package repositories

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/jmoiron/sqlx"
)

type AlbumsRepository interface {
	CreateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error)
	UploadImagesToAlbum(ctx context.Context, id string, uploads []models.UploadImages) ([]models.UploadPictureResponse, error)
	GetAlbum(ctx context.Context, id string) (models.AlbumWithImages, error)
	GetAlbums(ctx context.Context, orderBy, sort, search string, limit int, offset int) ([]models.AlbumsWithTotalCount, error)
	UpdateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error)
	DeleteAlbum(ctx context.Context, id string) (int, error)
	DeleteAlbumImage(ctx context.Context, key, id string) error
	SetAlbumCover(ctx context.Context, id string, imageName string) error
}

type albumsRepository struct {
	db            *sqlx.DB
	objectStorage *s3.Client
	Bucket        string
}

func NewAlbumsRepository(db *sqlx.DB, objectStorage *s3.Client) AlbumsRepository {
	return &albumsRepository{
		db:            db,
		objectStorage: objectStorage,
		Bucket:        internal.BUCKET_NAME,
	}
}

func (ar *albumsRepository) CreateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error) {
	return db.AddOneRow(
		ar.db,
		"./db/albums/post_album.sql",
		body,
	)
}

func (ar *albumsRepository) UploadImagesToAlbum(ctx context.Context, id string, uploads []models.UploadImages) ([]models.UploadPictureResponse, error) {
	var responses []models.UploadPictureResponse
	presignedClient := s3.NewPresignClient(ar.objectStorage)

	if !strings.HasSuffix(id, "/") {
		id += "/"
	}

	for _, upload := range uploads {

		randomBytes := make([]byte, 6)
		_, err := rand.Read(randomBytes)
		if err != nil {
			return nil, err
		}

		hash := sha256.Sum256(randomBytes)

		key := internal.ALBUM_PATH + id + "img_" + hex.EncodeToString(hash[:4]) + "_raw_" + upload.Filename

		presigned, err := presignedClient.PresignPutObject(ctx, &s3.PutObjectInput{
			Bucket:      aws.String(ar.Bucket),
			Key:         aws.String(key),
			ContentType: aws.String(upload.Type),
			ACL:         types.ObjectCannedACLPublicRead,
		}, s3.WithPresignExpires(10*time.Minute))
		if err != nil {
			return nil, err
		}

		responses = append(responses, models.UploadPictureResponse{
			URL:     presigned.URL,
			Headers: presigned.SignedHeader,
			Key:     key,
		})
	}

	return responses, nil
}

func (ar *albumsRepository) GetAlbum(ctx context.Context, id string) (models.AlbumWithImages, error) {
	album, err := db.ExecuteOneRow[models.AlbumWithImages](ar.db, "./db/albums/get_album.sql", id)
	if err != nil {
		return models.AlbumWithImages{}, err
	}

	path := id

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	prefix := internal.ALBUM_PATH + path

	var images []string
	paginator := s3.NewListObjectsV2Paginator(ar.objectStorage, &s3.ListObjectsV2Input{
		Bucket: aws.String(ar.Bucket),
		Prefix: aws.String(prefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return models.AlbumWithImages{}, err
		}

		for _, obj := range page.Contents {
			if strings.HasSuffix(*obj.Key, "/") {
				continue
			}

			images = append(images, strings.TrimPrefix(*obj.Key, prefix))
		}
	}

	album.Images = images
	album.ImageCount = len(images)

	return album, nil
}

func (ar *albumsRepository) GetAlbums(ctx context.Context, orderBy, sort, search string, limit int, offset int) ([]models.AlbumsWithTotalCount, error) {
	albums, err := db.FetchAllElements[models.AlbumsWithTotalCount](
		ar.db,
		"./db/albums/get_albums.sql",
		orderBy, sort,
		limit, offset,
		search,
	)
	if err != nil {
		return nil, err
	}

	images := make(map[string][]string)
	neededAlbums := make(map[string]bool)
	for _, album := range albums {
		albumID := strconv.Itoa(album.ID)
		neededAlbums[albumID] = true
	}

	paginator := s3.NewListObjectsV2Paginator(ar.objectStorage, &s3.ListObjectsV2Input{
		Bucket: aws.String(ar.Bucket),
		Prefix: aws.String(internal.ALBUM_PATH),
	})

	counts := make(map[string]int)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, obj := range page.Contents {
			if strings.HasSuffix(*obj.Key, "/") {
				continue
			}

			parts := strings.Split(*obj.Key, "/")
			if len(parts) >= 2 {
				albumID := parts[1]

				counts[albumID]++

				if !neededAlbums[albumID] {
					continue
				}

				if len(images[albumID]) < 3 {
					filename := parts[len(parts)-1]
					images[albumID] = append(images[albumID], filename)

					if len(images[albumID]) >= 3 {
						delete(neededAlbums, albumID)
					}
				}
			}
		}
	}
	for i := range albums {
		albumID := strconv.Itoa(albums[i].ID)
		albums[i].Images = images[albumID]
		albums[i].ImageCount = counts[albumID]
	}

	return albums, nil
}

func (ar *albumsRepository) UpdateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error) {
	return db.AddOneRow(
		ar.db,
		"./db/albums/put_album.sql",
		body,
	)
}

func (ar *albumsRepository) DeleteAlbum(ctx context.Context, id string) (int, error) {
	returnID, err := db.ExecuteOneRow[int](
		ar.db,
		"./db/albums/delete_album.sql",
		id,
	)

	var continuationToken *string

	prefix := internal.ALBUM_PATH + id
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	for {
		listOutput, err := ar.objectStorage.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
			Bucket:            aws.String(ar.Bucket),
			Prefix:            aws.String(prefix),
			ContinuationToken: continuationToken,
		})
		if err != nil {
			return 0, err
		}

		if len(listOutput.Contents) == 0 {
			break
		}

		var objectsToDelete []types.ObjectIdentifier
		for _, obj := range listOutput.Contents {
			objectsToDelete = append(objectsToDelete, types.ObjectIdentifier{
				Key: obj.Key,
			})
		}

		_, err = ar.objectStorage.DeleteObjects(ctx, &s3.DeleteObjectsInput{
			Bucket: aws.String(ar.Bucket),
			Delete: &types.Delete{
				Objects: objectsToDelete,
				Quiet:   aws.Bool(true),
			},
		})
		if err != nil {
			return 0, err
		}

		if *listOutput.IsTruncated {
			continuationToken = listOutput.NextContinuationToken
		} else {
			break
		}
	}
	return returnID, err
}

func (ar *albumsRepository) DeleteAlbumImage(ctx context.Context, key, id string) error {
	_, err := db.ExecuteOneRow[models.AlbumWithImages](
		ar.db,
		"./db/albums/get_album.sql",
		id,
	)
	if err != nil {
		return err
	}

	_, err = ar.objectStorage.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(ar.Bucket),
		Key:    aws.String(key),
	})
	return err
}

func (ar *albumsRepository) SetAlbumCover(ctx context.Context, id string, imageName string) error {
	prefix := internal.ALBUM_PATH + id + "/"

	listOutput, err := ar.objectStorage.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:  aws.String(ar.Bucket),
		Prefix:  aws.String(prefix + "coverimg_"),
		MaxKeys: aws.Int32(1),
	})
	if err != nil {
		return err
	}

	if len(listOutput.Contents) > 0 {
		firstKey := *listOutput.Contents[0].Key
		if strings.Contains(firstKey, "coverimg_") {
			oldCoverName := strings.Replace(firstKey, "coverimg_", "img_", 1)

			_, err = ar.objectStorage.CopyObject(ctx, &s3.CopyObjectInput{
				Bucket:     aws.String(ar.Bucket),
				CopySource: aws.String(ar.Bucket + "/" + firstKey),
				Key:        aws.String(oldCoverName),
			})
			if err != nil {
				return err
			}

			_, err = ar.objectStorage.DeleteObject(ctx, &s3.DeleteObjectInput{
				Bucket: aws.String(ar.Bucket),
				Key:    aws.String(firstKey),
			})
			if err != nil {
				return err
			}
		}
	}
	coverImageName := strings.Replace(imageName, "img_", "coverimg_", 1)
	path := internal.ALBUM_PATH + id + "/" + imageName
	coverPath := internal.ALBUM_PATH + id + "/" + coverImageName

	_, err = ar.objectStorage.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(ar.Bucket),
		CopySource: aws.String(ar.Bucket + "/" + path),
		Key:        aws.String(coverPath),
	})
	if err != nil {
		return err
	}

	_, err = ar.objectStorage.PutObjectAcl(ctx, &s3.PutObjectAclInput{
		Bucket: aws.String(ar.Bucket),
		Key:    aws.String(coverPath),
		ACL:    types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return err
	}

	_, err = ar.objectStorage.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(ar.Bucket),
		Key:    aws.String(internal.ALBUM_PATH + id + "/" + imageName),
	})
	return err
}

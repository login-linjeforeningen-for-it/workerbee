package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"slices"
	"strings"
	"workerbee/internal"
	"workerbee/repositories"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

var validPaths = []string{
	"events",
	"jobs",
	"organizations",
}

type ImageService struct {
	repo repositories.ImageRepository
}

func NewImageService(repo repositories.ImageRepository) *ImageService {
	return &ImageService{
		repo: repo,
	}
}

func (is *ImageService) UploadStorageProof(ctx context.Context) (string, error) {
	const key = internal.IMG_PATH + "events/workerbee-rustfs-proof.png"
	const image = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO+/p9sAAAAASUVORK5CYII="

	payload, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		return "", err
	}

	if err := is.repo.UploadImage(ctx, key, "image/png", bytes.NewReader(payload)); err != nil {
		return "", err
	}

	return key, nil
}

func (is *ImageService) UploadImage(file *multipart.FileHeader, ctx context.Context, path string) (string, error) {
	if !slices.Contains(validPaths, path) {
		return "", internal.ErrInvalidImagePath
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		return "", err
	}

	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	newW, newH := internal.DownscaleImage(w, h)
	if newW != w || newH != h {
		img = imaging.Resize(img, newW, newH, imaging.Box)
	}

	ops := &webp.Options{
		Lossless: false,
		Quality:  90,
	}

	buf := new(bytes.Buffer)
	if err := webp.Encode(buf, img, ops); err != nil {
		return "", err
	}

	key := internal.IMG_PATH + path + strings.Split(file.Filename, ".")[0] + ".webp"

	err = is.repo.UploadImage(ctx, key, "image/webp", buf)
	if err != nil {
		return "", err
	}

	return key, nil
}

func (is *ImageService) GetImagesInPath(ctx context.Context, path string) ([]string, error) {
	if !slices.Contains(validPaths, path) {
		return nil, internal.ErrInvalidImagePath
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	prefix := internal.IMG_PATH + path

	images, err := is.repo.GetImagesInPath(ctx, prefix)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (is *ImageService) DeleteImage(ctx context.Context, path, imageName string) (string, error) {
	if !slices.Contains(validPaths, path) {
		return "", internal.ErrInvalidImagePath
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	key := internal.IMG_PATH + path + imageName

	err := is.repo.DeleteImage(ctx, key)
	if err != nil {
		return "", err
	}

	return key, nil
}

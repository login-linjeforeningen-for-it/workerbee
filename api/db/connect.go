package db

import (
	"log"
	"workerbee/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init() *sqlx.DB {
	DB, err := sqlx.Connect("postgres", config.DB_url)
	if err != nil {
		log.Fatalln("Unable to connect to db: ", err)
	}
	log.Println("Connected to DB")
	return DB
}

func StorageInit() *s3.Client {
	client := s3.New(s3.Options{
		Region: config.StorageRegion,
		Credentials: aws.NewCredentialsCache(
			credentials.StaticCredentialsProvider{
				Value: aws.Credentials{
					AccessKeyID:     config.StorageAccessKeyID,
					SecretAccessKey: config.StorageSecretAccessKey,
				},
			},
		),
		EndpointResolver: s3.EndpointResolverFromURL(config.StorageURL),
		UsePathStyle:     true,
	})

	if client == nil {
		log.Fatalln("Unable to initialize S3-compatible object storage client")
	}
	log.Println("Initialized S3-compatible object storage client")
	return client
}

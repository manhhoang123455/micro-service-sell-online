package utils

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"mime/multipart"
	"product-service/config"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinio() {
	minioClient, err := minio.New(config.AppConfig.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AppConfig.MinioAccessKey, config.AppConfig.MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("failed to initialize MinIO client: %v", err)
	}

	MinioClient = minioClient

	err = MinioClient.MakeBucket(context.Background(), config.AppConfig.MinioBucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := MinioClient.BucketExists(context.Background(), config.AppConfig.MinioBucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", config.AppConfig.MinioBucketName)
		} else {
			log.Fatalf("failed to create bucket: %v", err)
		}
	} else {
		log.Printf("Successfully created %s\n", config.AppConfig.MinioBucketName)
	}
}

func UploadFile(file *multipart.FileHeader) (string, error) {
	fileName := file.Filename
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	_, err = MinioClient.PutObject(context.Background(), config.AppConfig.MinioBucketName, fileName, src, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", config.AppConfig.MinioEndpoint, config.AppConfig.MinioBucketName, fileName), nil
}

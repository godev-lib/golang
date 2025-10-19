package minio_custom

import (
	"log"

	"github.com/godev-lib/golang/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient(config *config.Config) *minio.Client {
	endpoint := config.Minio.Endpoint
	accessKeyID := config.Minio.AccessKeyID
	secretAccessKey := config.Minio.SecretAccessKey
	useSSL := config.Minio.UseSSL

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}

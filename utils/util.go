package utils

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/tech-rounak/jwt-authentication/service"
)

func SaveFile(fileHeader *multipart.FileHeader, key int) (string, error) {
	// fmt.Println(fileHeader.Filename)
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	cfg := service.LoadAWSConfig()
	s3Client := s3.NewFromConfig(*cfg)
	bucketName := "blinkit-task1"
	bucketKey := fileHeader.Filename
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &bucketKey,
		Body:   src,
	})

	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, bucketKey)
	// fmt.Println("url :: ", url)

	return url, nil
}

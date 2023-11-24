package util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DeleteFile(filename string) (err error) {
	sdkConfig, _ := config.LoadDefaultConfig(context.TODO())
	s3Client := s3.NewFromConfig(sdkConfig)

	_, _ = s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String("orders-confirmar"),
		Key:    aws.String(filename),
	})
	return err
}

package util

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"my-orders/cfg"
)

func UploadFile(file *multipart.File, cfg *cfg.Config, filename string) (location string, err error) {
	credential := credentials.NewStaticCredentialsProvider(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, "")
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credential), config.WithRegion(cfg.AWSRegion))

	if err != nil {
		return "", err
	}

	s3Client := s3.NewFromConfig(sdkConfig)

	uploader := manager.NewUploader(s3Client)

	up, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		BucketKeyEnabled: true,
		ACL:              types.ObjectCannedACLPublicRead,
		Bucket:           aws.String(cfg.AWSBucketName),
		Key:              aws.String(filename),
		Body:             *file,
	})

	if err != nil {
		return "", err
	}
	return up.Location, err
}

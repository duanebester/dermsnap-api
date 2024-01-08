package services

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ImageService interface {
	UploadImage(ctx context.Context, fileKey string, contentType string, reader io.Reader) (string, error)
}

type ImageServiceImpl struct {
	client     *s3.Client
	bucketName string
	pubBaseUrl string
}

func NewImageService(bucketName string) ImageService {
	var ctx = context.Background()
	var accountId = os.Getenv("R2_ACCOUNT_ID")
	var accessKeyId = os.Getenv("R2_ACCESS_KEY_ID")
	var accessKeySecret = os.Getenv("R2_ACCESS_KEY_SECRET")

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		panic("unable to load R2 SDK config, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)
	return ImageServiceImpl{
		client:     client,
		bucketName: bucketName,
		pubBaseUrl: os.Getenv("R2_PUBLIC_BASE_URL"),
	}
}

func (i ImageServiceImpl) UploadImage(ctx context.Context, fileKey string, contentType string, reader io.Reader) (string, error) {
	_, err := i.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &i.bucketName,
		Key:         &fileKey,
		Body:        reader,
		ContentType: &contentType,
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", i.pubBaseUrl, fileKey), nil
}

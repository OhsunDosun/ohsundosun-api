package s3

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type S3Info struct {
	Client     *s3.Client
	BucketName string
}

var Info S3Info

func init() {
	var bucketName = os.Getenv("R2_BUCKET_NAME")
	var accountId = os.Getenv("R2_ACCOUNT_ID")
	var accessKeyId = os.Getenv("R2_ACCESS_KEY_ID")
	var accessKeySecret = os.Getenv("R2_ACCESS_KEY_SECRET")

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	Info = S3Info{
		Client:     client,
		BucketName: bucketName,
	}
}

func (info S3Info) UploadFile(file multipart.FileHeader, key string) (*string, *error) {
	opendFile, err := file.Open()
	if err != nil {
		return nil, &err
	}

	fileNames := strings.Split(file.Filename, ".")
	fileName := uuid.NewString()

	if len(fileNames) > 0 {
		extension := fileNames[len(fileNames)-1]
		fileName += "." + extension
	}

	_, err = info.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(info.BucketName),
		Key:         aws.String(key + "/" + fileName),
		Body:        opendFile,
		ContentType: aws.String(file.Header.Get("Content-Type")),
	})
	if err != nil {
		return nil, &err
	}

	url := os.Getenv("IMAGE_HOST") + "/" + key + "/" + fileName

	return &url, nil
}

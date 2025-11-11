package awsclient

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type BucketBasics struct {
	S3Client *s3.Client
}

func InitS3Client() *BucketBasics {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := s3.NewFromConfig(cfg)
	return &BucketBasics{S3Client: client}
}

func (basics BucketBasics) UploadStream(
	ctx context.Context,
	bucketName string,
	objectKey string,
	body io.Reader,
) error {
	_, err := basics.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   body,
		ACL:    types.ObjectCannedACLPrivate, // or PublicRead if needed
	})
	return err
}

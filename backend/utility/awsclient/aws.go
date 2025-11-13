package awsclient

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

type BucketBasics struct {
	Client     *s3.Client
	BucketName string
	Region     string
}

// ─── Initialize AWS Client ───
func NewBucketBasics() (*BucketBasics, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return &BucketBasics{
		Client:     client,
		BucketName: os.Getenv("AWS_BUCKET_NAME"),
		Region:     os.Getenv("AWS_REGION"),
	}, nil
}

// ─── Upload Large Object ───
func (svc *BucketBasics) UploadLargeFile(ctx context.Context, key string, fileBytes []byte) (string, error) {
	uploader := manager.NewUploader(svc.Client, func(u *manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024 // 10MB per part
	})
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(svc.BucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(fileBytes),
		ACL:    types.ObjectCannedACLPrivate,
	})
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "EntityTooLarge" {
			log.Printf("Error: file too large for multipart upload to %s", svc.BucketName)
		} else {
			log.Printf("Couldn't upload large object to %v:%v. Error: %v\n", svc.BucketName, key, err)
		}
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", svc.BucketName, svc.Region, key)
	return url, nil
}

// ─── Download File ───
func (svc *BucketBasics) DownloadFile(ctx context.Context, key string) ([]byte, error) {
	var partMiBs int64 = 10
	downloader := manager.NewDownloader(svc.Client, func(d *manager.Downloader) {
		d.PartSize = partMiBs * 1024 * 1024
	})

	buffer := manager.NewWriteAtBuffer([]byte{})

	_, err := downloader.Download(ctx, buffer, &s3.GetObjectInput{
		Bucket: aws.String(svc.BucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		log.Printf("Couldn't download large object from %v:%v. Here's why: %v\n",
			svc.BucketName, key, err)
	}

	return buffer.Bytes(), err

}

// ─── List Files in Bucket ───
func (svc *BucketBasics) ListFiles(ctx context.Context, prefix string) ([]string, error) {
	resp, err := svc.Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(svc.BucketName),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}

	files := []string{}
	for _, item := range resp.Contents {
		files = append(files, *item.Key)
	}
	return files, nil
}

// ─── Delete File ───
func (svc *BucketBasics) DeleteFile(ctx context.Context, key string) error {
	_, err := svc.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(svc.BucketName),
		Key:    aws.String(key),
	})
	return err
}

package filehandlers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

// Download a file which is Read in CRUD for other REST API
func DownloadFile(c *gin.Context) {

	file, err := os.Create("pass_filename_here")
	if err != nil {
		log.Fatal("unable to create a file, ")
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Failed to open file",
		})
	}

	// STEP 1
	// Setup S3 Uploader. Load the Shared AWS Configuration (~/.aws/config)
	cfg, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	// STEP 2
	// Create an Amazon S3 service client
	client = s3.NewFromConfig(cfg)
	downloader := manager.NewDownloader(client)

	_, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String("key"),
	})
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"file": "file",
	})
}

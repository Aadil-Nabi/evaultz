package filehandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func DeleteFile(c *gin.Context) {
	// STEP 1
	// Setup S3 Uploader. Load the Shared AWS Configuration (~/.aws/config)
	cfg, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	// STEP 2
	// Create an Amazon S3 service client
	client = s3.NewFromConfig(cfg)
	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String("key"),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "file deleted successfully",
	})
}

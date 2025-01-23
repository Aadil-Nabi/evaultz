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

func ListFiles(c *gin.Context) {

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

	resp, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(BucketName),
	})
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"error": "Failed to load file from the form",
		})
	}
	for _, item := range resp.Contents {
		log.Printf("- %s (size: %d bytes)\n", aws.ToString(item.Key), item.Size)
	}

	c.JSON(http.StatusOK, gin.H{
		"output": resp,
	})

}

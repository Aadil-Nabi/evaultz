package filehandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

type deleteBody struct {
	Key      string
	Location string
}

func DeleteFile(c *gin.Context) {

	var deleteBody deleteBody

	c.Bind(&deleteBody)

	deleteBodyPayload := models.Files{
		Key: deleteBody.Key,
	}
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
		Key:    aws.String(deleteBodyPayload.Key),
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

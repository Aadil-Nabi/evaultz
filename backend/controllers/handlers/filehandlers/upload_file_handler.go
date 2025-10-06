package filehandlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

var cfg aws.Config
var err error
var client *s3.Client
var BucketName = "bucket.aadilnabi"

type FileDetails struct {
	Key      string
	Location string
}

var fileDetails FileDetails

func UploadFile(c *gin.Context) {

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

	// STEP 3
	// NewUploader creates a new Uploader instance to upload objects to S3
	uploader := manager.NewUploader(client)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Failed to load file from the form",
		})
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Failed to open file",
		})
	}

	// Upload file to S3 Bucket
	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("bucket.aadilnabi"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if uploadErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "unable to upload file",
		})
		return
	}

	Key := *result.Key
	Location := result.Location

	fmt.Println("KEY IS: ", Key)
	fmt.Println("LOCATION IS: ", Location)

	filePayload := models.Files{
		Key:      Key,
		Location: Location,
	}

	// Store the Key(file name) and URL(Location) of a file in a struct accessible outside from this package
	fileDetails := FileDetails{
		Key:      Key,
		Location: Location,
	}

	fileOnDB := configs.DB.Create(&filePayload)
	if fileOnDB.Error != nil {
		c.Status(400)
		return
	}

	fmt.Println("File stored is : ", fileOnDB)

	// Render the json
	c.JSON(http.StatusOK, gin.H{
		"result": fileDetails,
	})
}

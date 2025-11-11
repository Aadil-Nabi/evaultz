package filehandlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Aadil-Nabi/evaultz/utility/awsclient"
	"github.com/gin-gonic/gin"

	"path/filepath"
)

var s3Client = awsclient.InitS3Client()

func Upload(c *gin.Context) {
	// Parse file from form-data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Open the uploaded file
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
		return
	}
	defer f.Close()

	// Define bucket and object key
	bucket := "evaultz-s3-bucket"
	objectKey := filepath.Base(file.Filename) // or generate UUID

	// Upload to S3

	err = s3Client.UploadStream(context.TODO(), bucket, objectKey, f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload: %v", err)})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": file.Filename,
		"bucket":   bucket,
		"key":      objectKey,
	})
}

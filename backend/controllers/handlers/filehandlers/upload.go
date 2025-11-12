package filehandlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Aadil-Nabi/evaultz/utility/awsclient"
	"github.com/gin-gonic/gin"
)

// POST /upload-large
func UploadHandler(c *gin.Context) {
	bucketService, err := awsclient.NewBucketBasics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file not provided"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to open file"})
		return
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}

	key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), file.Filename)
	url, err := bucketService.UploadLargeFile(context.TODO(), key, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "upload failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "file uploaded successfully",
		"url":     url,
		"key":     key,
	})
}

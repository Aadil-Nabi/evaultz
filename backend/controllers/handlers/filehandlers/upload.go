package filehandlers

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/Aadil-Nabi/evaultz/utility/awsclient"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileUploadPayload struct {
	Visibility string `form:"visibility" json:"visibility"` // "public" or "private"
}

func UploadHandler(c *gin.Context) {
	db := configs.DB

	userID := c.MustGet("userID").(uuid.UUID)
	tenantID := c.MustGet("tenantID").(uuid.UUID)

	// Read file
	src, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	// Default visibility = private
	visibility := "private"

	var payload FileUploadPayload
	_ = c.ShouldBind(&payload)

	// Validate visibility value
	if payload.Visibility != "" {
		if payload.Visibility != "public" && payload.Visibility != "private" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "visibility must be 'public' or 'private'"})
			return
		}
		visibility = payload.Visibility
	}

	// Read file bytes
	fileReader, err := src.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer fileReader.Close()

	bytes, err := io.ReadAll(fileReader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}

	// Upload to S3
	s3, err := awsclient.NewBucketBasics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "s3 client init failed"})
		return
	}

	key := fmt.Sprintf("uploads/%s_%s", uuid.NewString(), src.Filename)

	url, err := s3.UploadLargeFile(context.TODO(), key, bytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed uploading to s3"})
		return
	}

	// Save DB record
	file := models.File{
		OwnerID:    userID,
		TenantID:   tenantID,
		FileName:   src.Filename,
		StorageKey: key,
		MimeType:   src.Header.Get("Content-Type"),
		Size:       src.Size,
		URL:        url,
		Visibility: visibility,
	}

	if err := db.Create(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed saving metadata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "file uploaded successfully",
		"file":    file,
		"url":     url, // This is just the S3 PUT return URL, not a public URL
	})
}

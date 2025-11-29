package filehandlers

import (
	"fmt"
	"net/http"

	"github.com/Aadil-Nabi/evaultz/utility/awsclient"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListHandler(c *gin.Context) {

	// Get tenantID, teamID, userID from the gin context.
	tenantID := c.MustGet("tenantID")
	userID := c.MustGet("userID")
	teamIDPtr, _ := c.Get("teamID")
	var teamID *uuid.UUID = nil

	if teamIDPtr != nil {
		if t, ok := teamIDPtr.(*uuid.UUID); ok {
			teamID = t
		}
	}

	s3, err := awsclient.NewBucketBasics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "s3 client init failed",
		})
	}

	prefix := fmt.Sprintf("uploads/%v/%v/%v/", tenantID, teamID, userID)

	files, err := s3.ListFiles(c, prefix)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"files": files,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})

}

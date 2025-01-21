package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCreateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "created",
	})
}

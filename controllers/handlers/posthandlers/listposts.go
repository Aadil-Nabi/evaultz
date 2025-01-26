package posthandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListPosts(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "this is the posts page",
	})
}

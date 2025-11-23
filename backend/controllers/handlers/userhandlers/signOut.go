package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	c.SetCookie(
		"jwt",
		"",
		-1, // expire immediately
		"/",
		"",
		true, // secure (HTTPS only)
		true, // httpOnly
	)
	c.JSON(http.StatusOK, gin.H{"message": " user logged out"})
}

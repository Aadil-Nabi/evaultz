package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("main function")

	// Create Gin Router.
	router := gin.Default()

	// Register Routes
	router.GET("/", homePage)
	router.GET("/user/:name", getUserName)
	// Start a Server.
	router.Run("localhost:9000")
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "This is my HomePage")
}

func getUserName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s ", name)
}

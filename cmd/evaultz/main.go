package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("main function")

	router := gin.Default()

	router.GET("/", homePage)
	router.Run("localhost:9000")
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "This is my HomePage")
}

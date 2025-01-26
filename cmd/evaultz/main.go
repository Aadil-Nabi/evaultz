package main

import (
	"fmt"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/controllers/handlers/posthandlers"
	"github.com/Aadil-Nabi/evaultz/controllers/handlers/userhandlers"
	"github.com/Aadil-Nabi/evaultz/middleware"
	"github.com/gin-gonic/gin"
)

// this is special function that runs before main and loads the encironment variables and configurations
func init() {
	configs.MustLoadEnvs()
	configs.ConnectDB()
}

func main() {

	fmt.Println("Hello from the main")

	// Create Gin Router.
	router := gin.Default()

	router.Use(gin.Logger())

	// Routers for user operations.
	router.POST("/signup", userhandlers.SignUpHandler)
	router.POST("/login", userhandlers.Login)

	// Routers for file operations
	// router.GET("/", filehandlers.ListFiles)
	// router.POST("/upload", filehandlers.UploadFile)
	// // router.GET("/download", controllers.DownloadFile)
	// router.DELETE("/delete", filehandlers.DeleteFile)

	// Routers for Posts
	router.GET("/posts", middleware.RequireAuth, posthandlers.ListPosts)

	router.Run()
}

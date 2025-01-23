package main

import (
	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/controllers/handlers/filehandlers"
	"github.com/Aadil-Nabi/evaultz/controllers/handlers/userhandlers"
	"github.com/gin-gonic/gin"
)

// this is special function that runs before main and loads the encironment variables and configurations
func init() {
	configs.MustLoadEnvs()
	configs.ConnectDB()
}

func main() {

	// Create Gin Router.
	router := gin.Default()

	// Routers for user operations.
	router.POST("/register", userhandlers.RegisterUserHandler)

	// Routers for file operations
	router.GET("/", filehandlers.ListFiles)
	router.POST("/upload", filehandlers.UploadFile)
	// router.GET("/download", controllers.DownloadFile)
	router.DELETE("/delete", filehandlers.DeleteFile)

	router.Run()
}

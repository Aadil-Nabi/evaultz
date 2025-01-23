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
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// router.GET("/", controllers.ListFiles)
	router.POST("/register", userhandlers.RegisterUserHandler)

	router.POST("/upload", filehandlers.UploadFile)
	// router.GET("/download", controllers.DownloadFile)

	// router.DELETE("/delete", controllers.DeleteFile)

	router.Run()
}

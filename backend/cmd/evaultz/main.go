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

	router.Use(gin.Logger())

	// Routers for user operations.
	router.POST("/api/v1/signup", userhandlers.SignUpHandler)
	router.POST("/api/v1/login", userhandlers.Login)
	router.PATCH("/api/v1/update/:id", userhandlers.UpdateUser)
	router.DELETE("/api/v1/delete/:id", userhandlers.DeleteUser)

	// Routes for user Detail.

	// Routers for file operations
	// router.GET("/", middleware.RequireAuth, filehandlers.ListFiles)
	router.POST("/api/v1/upload", filehandlers.UploadHandler)

	// Routers for Posts

	// Routes for Payments

	router.Run(":8082")
}

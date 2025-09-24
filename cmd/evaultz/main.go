package main

import (
	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/controllers/handlers/filehandlers"
	"github.com/Aadil-Nabi/evaultz/controllers/handlers/paymenthandlers"
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

	// Create Gin Router.
	router := gin.Default()

	router.Use(gin.Logger())

	// Routers for user operations.
	router.POST("/signup", userhandlers.SignUpHandler)
	router.POST("/login", userhandlers.Login)
	router.PATCH("/update/:id", userhandlers.UpdateUser)
	router.DELETE("/delete/:id", userhandlers.DeleteUser)

	// Routers for file operations
	router.GET("/", middleware.RequireAuth, filehandlers.ListFiles)
	router.POST("/upload", middleware.RequireAuth, filehandlers.UploadFile)
	// router.GET("/download", controllers.DownloadFile)
	router.DELETE("/delete", middleware.RequireAuth, filehandlers.DeleteFile)

	// Routers for Posts
	router.GET("/posts", middleware.RequireAuth, posthandlers.ListPosts)

	// Routes for Payments
	router.POST("/addpayment", middleware.RequireAuth, paymenthandlers.PaymentHandler)

	router.Run()
}

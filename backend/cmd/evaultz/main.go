package main

import (
	"net/http"
	"time"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/controllers/handlers/filehandlers"
	"github.com/Aadil-Nabi/evaultz/controllers/handlers/userhandlers"
	"github.com/Aadil-Nabi/evaultz/middleware"
	"github.com/Aadil-Nabi/evaultz/migrate"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// this is special function that runs before main and loads the encironment variables and configurations
func init() {
	configs.MustLoadEnvs()
	configs.ConnectDB()

	// Migrate daatbase/schema if not done, if done GORM automigrate will skip it.
	migrate.MigrateTables()
}

func main() {

	// Create Gin Router.
	router := gin.Default()

	router.Use(gin.Logger())

	//CORS, we can add it in  the middlewar, but for simplicity I added it here for time being.
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://frontend:3000"}, // frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routers for user operations.
	router.POST("/api/v1/signup", userhandlers.SignUpHandler)
	router.POST("/api/v1/signin", userhandlers.SignIn)
	router.POST("/api/v1/signout", userhandlers.SignOut)
	router.POST("/api/v1/forgotpassword", middleware.RequireAuth, userhandlers.ForgotPassword)
	// router.GET("/api/v1/me", middleware.RequireAuth, userhandlers.Me)
	// router.PATCH("/api/v1/update/:id", userhandlers.UpdateUser)
	router.DELETE("/api/v1/delete", middleware.RequireAuth, userhandlers.DeleteUser)

	// Routes for user Detail.
	router.GET("/api/v1/me", middleware.RequireAuth, func(c *gin.Context) {
		user := c.MustGet("user")

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	// Routers for file operations
	// router.GET("/", middleware.RequireAuth, filehandlers.UploadHandler)
	router.POST("/api/v1/upload", middleware.RequireAuth, filehandlers.UploadHandler)
	router.GET("/api/v1/download/:filename", middleware.RequireAuth, filehandlers.DownloadHander)
	router.GET("/api/v1/list", middleware.RequireAuth, filehandlers.ListHandler)

	// Routers for Posts

	// Routes for Payments

	router.Run(":8082")
}

package userhandlers

import (
	"net/http"
	"os"
	"time"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userLoginDetails struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	CompanyName string `json:"companyname" binding:"required"`
}

func SignIn(c *gin.Context) {
	var loginDetails userLoginDetails

	// Validate request
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Uniform error response for security
	invalidCreds := gin.H{"error": "Invalid credentials"}

	// DB lookup
	var user models.User
	err := configs.DB.
		Where("email = ? AND company_name = ?", loginDetails.Email, loginDetails.CompanyName).
		First(&user).Error

	if err != nil {
		// Avoid revealing whether email exists
		c.JSON(http.StatusUnauthorized, invalidCreds)
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, invalidCreds)
		return
	}

	// Validate secret
	secret := os.Getenv("TOKEN_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server misconfigured"})
		return
	}

	// JWT creation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     user.ID.String(),
		"email":   user.Email,
		"company": user.CompanyName,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Cookie security options
	secure := os.Getenv("ENV") == "prod"

	c.SetSameSite(http.SameSiteLaxMode)
	if secure {
		c.SetSameSite(http.SameSiteStrictMode)
	}

	c.SetCookie(
		"jwt",
		tokenString,
		3600*24, // 24 hours
		"/",
		"",
		secure, // secure only in prod
		true,   // HttpOnly
	)

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"message":  "Login successful",
	})
}

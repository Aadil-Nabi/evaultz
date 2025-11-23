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
	Email       string `json:"email"`
	Password    string `json:"password"`
	CompanyName string `json:"companyname"`
}

func SignIn(c *gin.Context) {
	var login userLoginDetails

	// Parse JSON body
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Lookup user
	var user models.User

	result := configs.DB.Where("email = ? AND company_name = ?", login.Email, login.CompanyName).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or company"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),                      // must be string
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 24h expiry
		"iat": time.Now().Unix(),                     // issued at
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign token"})
		return
	}

	// Set cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"jwt",       // cookie name
		tokenString, // value
		3600*24,     // expires in 24 hours
		"/",         // path
		"",          // domain (set in prod)
		true,        // secure (HTTPS only in prod)
		true,        // HTTPOnly
	)

	// Response (WITHOUT token — cookies only)
	c.JSON(http.StatusOK, gin.H{
		"id":      user.ID,
		"message": user.Username + " logged in successfully",
	})
}

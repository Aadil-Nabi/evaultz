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
	Email    string
	Password string
}

func Login(c *gin.Context) {
	var userLoginDetails userLoginDetails

	// Get the user details from the request.
	if c.Bind(&userLoginDetails) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to get the email and password for the user",
		})
		return
	}

	// Lookup the requested user in the DB ans store in a variable
	var user models.User

	configs.DB.First(&user, "email=?", userLoginDetails.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare the password hash with the one store in DB
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginDetails.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed Login, Invalid email and password",
		})
	}

	// Generate JWT token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		// "expiry": time.Now().Add(time.Second * 60).Unix(),
		"expiry": time.Now().Add(time.Hour * 1).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
	}

	// Store this jwt token in a cookie instead of sending it on JSON payload
	c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600*24, "", "", true, true)
	c.SetCookie("Authorization", tokenString, 3600, "", "", true, true)

	// Send a response.
	c.JSON(http.StatusOK, gin.H{
		"id":      user.ID,
		"message": user.FirstName + " logged in successfully",
	})

}

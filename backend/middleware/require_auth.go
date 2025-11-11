package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off the request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// TOKEN_SECRET is a []byte containing the secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check the expiry
		if float64(time.Now().Unix()) > claims["expiry"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Try to Login or Register",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token sub
		var user models.User
		configs.DB.First(&user, claims["sub"])
		if user.ID == [16]byte{} {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to request
		c.Set("user", user)

		// Continue to next
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)

	}

}

package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func RequireAuth(c *gin.Context) {
	// 1. Read secure cookie
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication"})
		c.Abort()
		return
	}

	// 2. Parse JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// 3. Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token structure"})
		c.Abort()
		return
	}

	// 4. Check expiry ("exp" not "expiry")
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			c.Abort()
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid expiration"})
		c.Abort()
		return
	}

	// 5. Fetch user
	userIDStr := claims["sub"].(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		c.Abort()
		return
	}

	// var user models.User
	// result := configs.DB.First(&user, "id = ?", userID)

	// if result.Error != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
	// 	c.Abort()
	// 	return
	// }

	fmt.Println("USER FROM MIDDLEWARE", userID)

	// 6. Attach user to context
	c.Set("userID", userID)

	c.Next()
}

package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func RequireAuth(c *gin.Context) {
	// 1️⃣ Read Cookie
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authentication"})
		c.Abort()
		return
	}

	// 2️⃣ Parse Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		c.Abort()
		return
	}

	// 3️⃣ Extract Claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token structure"})
		c.Abort()
		return
	}

	// 4️⃣ Validate Expiry
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
			c.Abort()
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing expiration"})
		c.Abort()
		return
	}

	// 5️⃣ Extract User ID
	userIDStr, ok := claims["sub"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing user id"})
		c.Abort()
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
		c.Abort()
		return
	}

	// 6️⃣ Extract Tenant ID
	tenantIDStr, ok := claims["tenant_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing tenant id"})
		c.Abort()
		return
	}

	tenantID, err := uuid.Parse(tenantIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid tenant id"})
		c.Abort()
		return
	}

	// 7️⃣ Extract Team ID (OPTIONAL - empty string means no team)
	var teamID *uuid.UUID = nil

	if v, exists := claims["team_id"]; exists {
		if idStr, ok := v.(string); ok && idStr != "" && idStr != "null" {
			id, err := uuid.Parse(idStr)
			if err == nil {
				teamID = &id
			}
		}
	}

	// 8️⃣ Extract Username (optional)
	username, _ := claims["username"].(string)

	// 9️⃣ Inject into Gin Context
	c.Set("userID", userID)
	c.Set("tenantID", tenantID)
	c.Set("teamID", teamID) // may be nil
	c.Set("username", username)

	// 🔟 Continue
	c.Next()
}

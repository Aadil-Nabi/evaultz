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

func SignIn(c *gin.Context) {
	type userLoginDetails struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Tenant   string `json:"tenant" binding:"required"`
	}

	var login userLoginDetails

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login payload"})
		return
	}

	invalid := gin.H{"error": "invalid credentials"}

	// 1️⃣ Find tenant
	var tenant models.Tenant
	if err := configs.DB.Where("name = ?", login.Tenant).First(&tenant).Error; err != nil {
		c.JSON(http.StatusUnauthorized, invalid)
		return
	}

	// 2️⃣ Find user inside tenant
	var user models.User
	err := configs.DB.
		Where("email = ? AND tenant_id = ?", login.Email, tenant.ID).
		Preload("Team").
		First(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, invalid)
		return
	}

	// 3️⃣ Validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, invalid)
		return
	}

	secret := os.Getenv("TOKEN_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server misconfigured"})
		return
	}

	// 4️⃣ Build JWT claims
	claims := jwt.MapClaims{
		"sub":       user.ID.String(),
		"email":     user.Email,
		"tenant_id": user.TenantID.String(),
		"tenant":    tenant.Name,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	// 5️⃣ Add team only if exists
	if user.TeamID != nil && user.Team != nil {
		claims["team_id"] = user.Team.ID.String()
		claims["team"] = user.Team.Name
	} else {
		claims["team_id"] = nil
		claims["team"] = nil
	}

	// 6️⃣ Sign JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// 7️⃣ Write JWT Cookie
	secure := os.Getenv("ENV") == "prod"
	c.SetSameSite(http.SameSiteLaxMode)

	if secure {
		c.SetSameSite(http.SameSiteStrictMode)
	}

	c.SetCookie(
		"jwt",
		tokenString,
		3600*24,
		"/",
		"",
		secure,
		true,
	)

	// 8️⃣ Response
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"tenant":   tenant.Name,
		"team":     claims["team"],
		"message":  "login successful",
	})
}

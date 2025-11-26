package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ForgotPasswordDetails struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ForgotPassword(c *gin.Context) {
	var body ForgotPasswordDetails

	// bind the input JSON data and store in a body  object of type ForgotPasswordDeatils struct defined above.
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to bind input payload",
		})
		return
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to hash the password",
		})
		return
	}

	// var user models.User

	// Store the hashed password in a database
	result := configs.DB.Model(&models.User{}).
		Where("email = ? AND username = ?", body.Email, body.Username).
		Update("password", string(hashedPassword))

	fmt.Println("Hashed Pass:", string(hashedPassword))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found or data mismatch",
		})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   result.Error.Error(),
			"message": "failed to update password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password updated successfully",
	})
}

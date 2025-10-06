package userhandlers

import (
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// user struct to store the values received from the request body
type userDetails struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	Phone     string
	DOB       string
}

func SignUpHandler(c *gin.Context) {
	var userbody userDetails

	// Bind decodes the json payload received on the request body into the struct specified as a pointer
	err := c.Bind(&userbody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(userbody.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to Hash password",
		})
		return
	}

	// initialize and assign the values received from the Jason Payload from user, to the User struct
	user := models.User{
		FirstName: userbody.FirstName,
		LastName:  userbody.LastName,
		Username:  userbody.Username,
		Email:     userbody.Email,
		Password:  string(hash),
		Phone:     userbody.Phone,
		DOB:       userbody.DOB,
	}

	// Create the user
	usr := configs.DB.Create(&user)
	if usr.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   usr.Error.Error(),
			"message": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})

	// result := configs.DB.Where("email = ?", user.Email).Find(&user)
	// if result.RowsAffected > 0 {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "email already exists",
	// 	})
	// }

}

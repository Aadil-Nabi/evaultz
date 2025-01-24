package userhandlers

import (
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
)

// user struct to store the values received from the request body
type userDetails struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func RegisterUserHandler(c *gin.Context) {
	var userbody userDetails
	//  Bind decodes the json payload received on the request body into the struct specified as a pointer
	c.Bind(&userbody)

	user := models.User{
		FirstName: userbody.FirstName,
		LastName:  userbody.LastName,
		Email:     userbody.Email,
		Password:  userbody.Password,
	}

	result := configs.DB.Where("email = ?", user.Email).Find(&user)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "email already exists",
		})
	} else {
		usr := configs.DB.Create(&user)
		if usr.Error != nil {
			c.Status(400)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": user,
		})
	}

}

package userhandlers

import (
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
)

// user struct to store the values received from the request body
type userBody struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func RegisterUserHandler(c *gin.Context) {
	var user userBody
	//  Bind decodes the json payload received on the request body into the struct specified as a pointer
	c.Bind(&user)

	userPayload := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	usr := configs.DB.Create(&userPayload)
	if usr.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": userPayload,
	})
}

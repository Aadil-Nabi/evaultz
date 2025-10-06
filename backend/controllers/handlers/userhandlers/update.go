package userhandlers

import (
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
)

// user struct to store the values received from the request body
type updatedUserDetails struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func UpdateUser(c *gin.Context) {
	// Get the ID of the user from the Param
	id := c.Param("id")

	// Get user from the database using ID of the user and store it in user variable of type models.User
	var user models.User
	configs.DB.First(&user, "id=?", id)

	// Get the fields to be updated from the Json Input.
	var updatedUser updatedUserDetails
	if err := c.Bind(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to bind the input to be updated",
		})
	}

	// Update multiple columns for the user details in the database.
	configs.DB.Model(&user).Updates(&updatedUser)
	c.JSON(http.StatusOK, gin.H{
		"result": updatedUser,
	})

}

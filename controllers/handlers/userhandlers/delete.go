package userhandlers

import (
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {

	// Get the ID of the user to be deleted
	id := c.Param("id")

	var user models.User
	result := configs.DB.Delete(&user, "id=?", id)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "unable to delete user",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user " + user.Email + "deleted successfully",
	})

}

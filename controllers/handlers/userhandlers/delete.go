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

	// configs.DB.Where("id=?", id).Delete(&user)
	configs.DB.Delete(&user, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "user " + user.Email + "deleted successfully",
	})

}

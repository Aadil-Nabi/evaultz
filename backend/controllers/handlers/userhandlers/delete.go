package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {

	// Get the ID of the user to be deleted
	userID, exists := c.Get("userID")

	var user models.User

	fmt.Println("User to be DELETED is", userID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not logged in",
		})
	}

	response := configs.DB.Delete(&user, userID)

	// if response.Error(c.JSON(http.StatusOK, gin.H{
	// 	"message": "unable to delete user",
	// }))

	if response.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to delete user",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "user with ID deleted successfully",
			"user": userID,
		})
	}

	// configs.DB.Where("id=?", id).Delete(&user)

}

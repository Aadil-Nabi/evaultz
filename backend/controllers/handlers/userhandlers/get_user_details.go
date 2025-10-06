package userhandlers

import (
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {

	var user models.User

	// get the user from the gin context which was set in the middleware.
	usr, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	user, ok := usr.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user in context"})
		return
	}

	configs.DB.Preload("Cards").First(&user, user.ID)
	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": "unauthorized",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	// for _, card := range user.Cards {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"result": user.Cards,
	// 	})
	// }

}

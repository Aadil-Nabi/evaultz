package userhandlers

import (
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
)

// type UserDetail struct {
// 	FirstName string
// 	LastName  string
// 	Username  string
// 	Email     string
// 	Password  string
// 	Phone     string
// 	DOB       string
// }

// var userDetail UserDetail

func GetUserDetails(c *gin.Context) {

	// // Bind the input payload to a struct and store the incoming values inside the struct created above.
	// err := c.Bind(&userDetail)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err,
	// 	})
	// }

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
		"user":  user,
		"cards": user.Cards,
	})
	// for _, card := range user.Cards {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"result": user.Cards,
	// 	})
	// }

}

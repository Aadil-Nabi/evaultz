package paymenthandlers

import (
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
)

type cardDetails struct {
	Name       string
	CardNumber string
	Cvv        string
	Expiry     string
}

func PaymentHandler(c *gin.Context) {

	// create and instantiate a variable of cardDetails struct to store the input values.
	var cardDetailsBody cardDetails

	err := c.Bind(&cardDetailsBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to bind/parse the input json payload",
		})
	}

	cardDetailModel := models.Card{
		Name:       cardDetailsBody.Name,
		CardNumber: cardDetailsBody.CardNumber,
		Cvv:        cardDetailsBody.Cvv,
		Expiry:     cardDetailsBody.Expiry,
	}

	cardDetail := configs.DB.Create(&cardDetailModel)

	if cardDetail.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to store card inside the database",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": cardDetail,
	})
}

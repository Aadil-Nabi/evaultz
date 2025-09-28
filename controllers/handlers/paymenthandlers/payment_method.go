package paymenthandlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	httpClient "github.com/Aadil-Nabi/evaultz/internal/pkg"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
)

type cardDetails struct {
	Name                   string
	CardNumber             string
	Cvv                    string
	Expiry                 string
	Protection_policy_name string
	Data                   string
}

type ProtectPolicy struct {
	Protection_policy_name string `json:"protection_policy_name"`
	Data                   string `json:"data"`
}

var cardDetailsBody cardDetails

func PaymentHandler(c *gin.Context) {

	// create and instantiate a variable of cardDetails struct to store the input values.

	err := c.Bind(&cardDetailsBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to bind/parse the input json payload",
		})
	}

	// Encrypt the credit card number
	encryptedData := encrypting()
	encryptedCreditCard := encryptedData["protected_data"]

	cardDetailModel := models.Card{
		Name:       cardDetailsBody.Name,
		CardNumber: encryptedCreditCard,
		Cvv:        cardDetailsBody.Cvv,
		Expiry:     cardDetailsBody.Expiry,
	}

	cardDetail := configs.DB.Create(&cardDetailModel)

	if cardDetail.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to store card inside the database",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": cardDetailModel,
	})

}

func encrypting() map[string]string {

	creditCardNumber := cardDetailsBody.CardNumber

	// url := cnfg.Base_Url + cnfg.Version + "/crypto/encrypt"
	url := "http://ciphertrust:8090/v1/protect"

	// Encode the data to be encrypted in base64 string as CM only accepts a valid base64 string
	plaintext := creditCardNumber
	plaintext = base64.StdEncoding.EncodeToString([]byte(plaintext))
	payload := map[string]string{
		"protection_policy_name": "internal_pp",
		"data":                   plaintext,
	}

	// Convert data into JSON encoded byte array
	encodedBody, _ := json.Marshal(payload)

	// convert the encoded JSON data to a type implemented by the io.Reader interface
	body := bytes.NewBuffer(encodedBody)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalf("Something went wrong in the request  %v", err)
	}

	// Add the required headers to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	//get client from a helper function
	client := httpClient.GetClient()

	// Do method to send the http request to the CM to http response
	// this is used when we add headers to the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to Encrypt %v", err)
	}

	// close the response
	defer resp.Body.Close()

	// Read the response received from the CM
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var output map[string]string

	// yaml.Unmarshal(data, &output)
	json.Unmarshal(data, &output)
	log.Println("protected credit card using CRDP ", output["protected_data"])

	return output

}

package main

import (
	"fmt"
	"log"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
)

func main() {
	migrateTenant()
}

func migrateTenant() {

	configs.MustLoadEnvs()
	configs.ConnectDB()

	// initialize and assign the values received from the Jason Payload from user, to the User struct
	// tenant := models.Tenant{
	// 	Name: "Thales",
	// }

	// Create the user inside the DB
	// t := configs.DB.Create(&tenant)
	// if t.Error != nil {
	// 	fmt.Println("tenant", t, "created")
	// }

	var tenantt models.Tenant

	configs.DB.Where("name = ?", "Thales").First(&tenantt)
	ID := tenantt.ID

	team := models.Team{
		TenantID: ID,
		Name:     "Engineering",
	}

	tm := configs.DB.Create(&team)
	if tm.Error != nil {
		fmt.Println("team", tm, "created")
	}

	log.Println("Tenant and Team initialized successfully")

}

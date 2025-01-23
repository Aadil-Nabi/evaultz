package main

import (
	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
)

// this is special function that runs before main and loads the encironment variables and configurations
func init() {
	configs.MustLoadEnvs()
	configs.ConnectDB()
}

// This function will only be used to create a Table for the first time and won't be
// used anyfurther once the table is created from the struct
func main() {

	// Call Automigrate on DB instaance received from the configs package
	// configs.DB.AutoMigrate(models.Post{})
	configs.DB.AutoMigrate(models.User{})
	configs.DB.AutoMigrate(models.Files{})
}

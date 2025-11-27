package migrate

import (
	"log"

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
// func main() {
func MigrateTables() {

	// Call Automigrate on DB instaance received from the configs package

	configs.DB.AutoMigrate(
		&models.Tenant{},
		&models.Team{},
		&models.User{},
		&models.File{},
	)

	log.Println("Successfully created tables inside the database")
}

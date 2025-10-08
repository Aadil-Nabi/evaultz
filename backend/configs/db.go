package configs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable so that we can access these variables anywhere outside of this package
var DB *gorm.DB
var err error

func ConnectDB() {

	dsn := os.Getenv("DB_URL")
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database : ", err)
	}

	fmt.Println("✔ Successfully Connected to the database", DB.Name())

}

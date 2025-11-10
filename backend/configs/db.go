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

	dsn := os.Getenv("DATABASE_URL")
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database : %v ", err)
	}

	fmt.Println("✔ Successfully Connected to the  Neon Postgres database", DB.Name())

}

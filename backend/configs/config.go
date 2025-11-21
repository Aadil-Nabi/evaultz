package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func MustLoadEnvs() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("unable to load .env file", err)
	}

}

package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func MustLoadEnvs() {

	err := godotenv.Load()
	if err != nil {
		log.Println("unable to load .env file", err)
	}

}

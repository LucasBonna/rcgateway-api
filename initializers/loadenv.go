package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port      string
	DbConnStr string
	EHGATEWAY string
	EHCRAWLER string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}

	Port = os.Getenv("PORT")
	DbConnStr = os.Getenv("DBCONNSTR")
	EHGATEWAY = os.Getenv("EHGATEWAY")
	EHCRAWLER = os.Getenv("EHCRAWLER")
}

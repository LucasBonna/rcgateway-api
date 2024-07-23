package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port            string
	DbConnStr       string
	RCGATEWAY       string
	RCAUTH          string
	RCSTORAGE       string
	RCTRACKER       string
	RCNOTIFICATIONS string
	RCREGISTRY      string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}

	Port = os.Getenv("PORT")
	DbConnStr = os.Getenv("DBCONNSTR")
	RCGATEWAY = os.Getenv("RCGATEWAY")
	RCAUTH = os.Getenv("RCAUTH")
	RCSTORAGE = os.Getenv("RCSTORAGE")
	RCTRACKER = os.Getenv("RCTRACKER")
	RCNOTIFICATIONS = os.Getenv("RCNOTIFICATIONS")
	RCREGISTRY = os.Getenv("RCREGISTRY")
}

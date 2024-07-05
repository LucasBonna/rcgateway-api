package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Port string
var Db_conn_str string
var RCGATEWAY string
var RCAUTH string
var RCSTORAGE string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}

	Port = os.Getenv("PORT")
	Db_conn_str = os.Getenv("DB_CONN_STR")
	RCGATEWAY=os.Getenv("RCGATEWAY")
	RCAUTH=os.Getenv("RCAUTH")
	RCSTORAGE=os.Getenv("RCSTORAGE")
}
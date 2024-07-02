package database

import (
	"log"
	"web/gin/initializers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDB() {
	var err error
	dsn := initializers.Db_conn_str
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to DB")
	} else {
		log.Println("Successfully connected to DB")
	}
}
package database

import (
	"log"
	"time"
	"web/gin/initializers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDB() {
	const maxRetries = 5
	const retryDelay = 2 * time.Second

	dsn := initializers.DbConnStr
	for tries := 1; tries <= maxRetries; tries++ {
		var err error
		Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to DB")
			return
		}

		log.Printf("Error connecting to DB (attempt %d/%d): %v", tries, maxRetries, err)
		if tries < maxRetries {
			log.Printf("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	log.Fatal("Failed to connect to DB after maximum retries")
}

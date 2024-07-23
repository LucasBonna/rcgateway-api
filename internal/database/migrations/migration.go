package migrations

import (
	"rc/gateway/internal/database"
	"rc/gateway/internal/database/models"
)

func MigrateModels() {
	database.Db.AutoMigrate(&models.Log{})
}


package migrations

import (
	"web/gin/internal/database"
	"web/gin/internal/database/models"
)

func MigrateModels() {
	database.Db.AutoMigrate(&models.Log{})
}
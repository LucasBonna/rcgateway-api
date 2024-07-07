package main

import (
	"web/gin/initializers"
	"web/gin/internal/database"
	"web/gin/internal/database/migrations"
	"web/gin/internal/server"
)

// @title Web App API
// @version 1.0
// @description Esta é a API da Web App.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	initializers.LoadEnv()
	database.ConnectToDB()
	migrations.MigrateModels()
	server.Start()
}
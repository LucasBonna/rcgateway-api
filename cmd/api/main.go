package main

import (
	"web/gin/initializers"
	"web/gin/internal/database"
	"web/gin/internal/database/migrations"
	"web/gin/internal/server"
)

func init() {

}

func main() {
	initializers.LoadEnv()
	database.ConnectToDB()
	migrations.MigrateModels()
	server.Start()
}
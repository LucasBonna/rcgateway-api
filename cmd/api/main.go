package main

import (
	"rc/gateway/initializers"
	"rc/gateway/internal/database"
	"rc/gateway/internal/database/migrations"
	"rc/gateway/internal/server"
)

// @title EHGateway App API
// @version 1.0
// @description Esta é a API da EHGateway App.

// @host localhost:8000

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	initializers.LoadEnv()
	database.ConnectToDB()
	migrations.MigrateModels()
	server.Start()
}

package main

import (
	"web/gin/initializers"
	"web/gin/internal/database"
	"web/gin/internal/server"
)

func main() {
	initializers.LoadEnv()
	database.ConnectToDB()
	server.Start()
}
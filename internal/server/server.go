package server

import (
	"web/gin/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowAllOrigins:  true,
        AllowMethods:     []string{"PUT", "PATCH", "DELETE", "POST", "GET", "HEAD"},
        AllowHeaders:     []string{"Origin"},
        AllowCredentials: true,
    }))

	registerRoutes(r)

	r.Run(":" + initializers.Port)
}
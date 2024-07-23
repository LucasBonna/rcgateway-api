package server

import (
	"rc/gateway/initializers"
	"rc/gateway/internal/server/routes"

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

	routes.RegisterRoutes(r)

	r.Run(":" + initializers.Port)
}


package routes

import (
	"web/gin/controllers"
	"web/gin/internal/middlewares"

	"github.com/gin-gonic/gin"
)


// @title Web App API
// @version 1.0
// @description Esta é a API da Web App.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func RegisterRoutes(r *gin.Engine) {
	r.Use(middlewares.Logger())

	r.GET("/swagger.json", controllers.JsonHandler)

	r.GET("/scalar", controllers.ScalarHandler)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controllers.PingHandler)
	}
}
package routes

import (
	"web/gin/controllers"
	"web/gin/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(middlewares.Logger())

	r.Use(middlewares.ReverseProxy())

	r.GET("/swagger.json", controllers.JsonHandler)

	r.GET("/scalar", controllers.ScalarHandler)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controllers.PingHandler)
	}
}
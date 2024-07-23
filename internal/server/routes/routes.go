package routes

import (
	"rc/gateway/controllers"
	"rc/gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(middlewares.Logger())

	r.Use(middlewares.ReverseProxy())

	r.GET("/docs", controllers.RedocHandler)

	r.GET("/merged-docs", controllers.MergedDocs)

	r.GET("/swagger.json", controllers.JsonHandler)

	r.GET("/swagger", controllers.SwaggerRedirect)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/merged-docs")))

	r.GET("/scalar", controllers.ScalarHandler)

	v1 := r.Group("/rcgateway")
	{
		v1.GET("/ping", controllers.PingHandler)
	}
}


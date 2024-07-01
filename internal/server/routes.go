package server

import (
	"log"
	"web/gin/internal/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func registerRoutes(r *gin.Engine) {
	r.Use(middlewares.Logger())

	r.GET("/ping", func(c *gin.Context) {
		example := c.MustGet("example").(uuid.UUID)

		log.Println(example)

		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
}
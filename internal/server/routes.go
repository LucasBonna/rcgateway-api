package server

import (
	"net/http"
	"web/gin/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Body string
}

func registerRoutes(r *gin.Engine) {
	r.Use(middlewares.Logger())

	r.POST("/post", func(c *gin.Context) {
		var json Post
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"error": "Bad Request",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H {
			"message": "Hello, " + json.Body,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
}
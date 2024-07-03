package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Body string `json:"body" binding:"required"`
}

func TesteRoutes(r *gin.Engine) {
	r.POST("/post", func(c *gin.Context) {
		var json Post
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"error": "Invalid request body",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H {
			"message": "Hello, " + json.Body,
		})
	})
}
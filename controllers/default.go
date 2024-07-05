package controllers

import "github.com/gin-gonic/gin"

// @Summary Ping
// @Description Verifica se o servidor está ativo
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H {
		"message": "pong",
	})
}
package server

import (
	"log"
	"net/http"
	"path/filepath"
	"web/gin/internal/middlewares"
	"web/gin/internal/server/routes"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)



func registerRoutes(r *gin.Engine) {
	r.Use(middlewares.Logger())

	routes.TesteRoutes(r)

	r.GET("/scalar", func(c *gin.Context) {
        specPath := filepath.Join("cmd", "docs", "swagger.json")
        htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
            SpecURL: specPath,
            CustomOptions: scalar.CustomOptions{
                PageTitle: "Simple API",
            },
            DarkMode: true,
        })

        if err != nil {
            log.Printf("%v", err)
            c.String(http.StatusInternalServerError, "Erro ao gerar a referência da API")
            return
        }

        c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
    })

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
}
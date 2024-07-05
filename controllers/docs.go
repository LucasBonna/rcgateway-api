package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)

func ScalarHandler(c *gin.Context) {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: "http://localhost:8080/swagger.json",
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
}

func JsonHandler(c *gin.Context) {
	data, err := os.ReadFile("./cmd/docs/swagger.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}
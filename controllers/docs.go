package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)

type SwaggerDoc struct {
	Swagger string `json:"swagger"`
	Info    struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Version     string `json:"version"`
	} `json:"info"`
	Host       string                 `json:"host"`
	BasePath   string                 `json:"basePath"`
	Paths      map[string]interface{} `json:"paths"`
	Components struct {
		Schemas map[string]interface{} `json:"schemas"`
	} `json:"components"`
	Tags        []interface{}          `json:"tags"`
	Definitions map[string]interface{} `json:"definitions"`
}

func MergedDocs(c *gin.Context) {
	localDocs := SwaggerDoc{
		Swagger: "2.0",
		Info: struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Version     string `json:"version"`
		}{
			Title:       "Receba! API",
			Description: "Combined API documentation",
			Version:     "1.0.0",
		},
		Host:  "localhost:3333",
		Paths: make(map[string]interface{}),
		Components: struct {
			Schemas map[string]interface{} `json:"schemas"`
		}{Schemas: make(map[string]interface{})},
		Tags:        []interface{}{},
		Definitions: make(map[string]interface{}),
	}

	// Load local Swagger JSON
	localSwagger, err := loadLocalSwagger()
	if err == nil {
		localDocs.Paths = localSwagger.Paths
		localDocs.Components.Schemas = localSwagger.Components.Schemas
		localDocs.Tags = localSwagger.Tags
		localDocs.Definitions = localSwagger.Definitions
	}

	services := []string{
		"http://localhost:3333/swagger.json",
		"http://rcauth/swagger.json",
		"http://rcstorage-api/swagger.json",
		"http://rctracker-api/swagger.json",
		"http://rcnotifications-api/swagger.json",
		"http://rcregistry-api/swagger.json",
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	client := &http.Client{Timeout: 5 * time.Second}
	done := make(chan bool)

	for _, url := range services {
		wg.Add(1)
		go func(url string) {
			defer func() {
				wg.Done()
				done <- true
			}()
			resp, err := client.Get(url)
			if err != nil {
				log.Printf("failed loading docs from %s: %v", url, err)
				return
			}
			defer resp.Body.Close()

			var data SwaggerDoc
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				log.Println(err)
				return
			}

			mu.Lock()
			if localDocs.Paths == nil {
				localDocs.Paths = make(map[string]interface{})
			}
			for k, v := range data.Paths {
				localDocs.Paths[k] = v
			}
			if localDocs.Components.Schemas == nil {
				localDocs.Components.Schemas = make(map[string]interface{})
			}
			if data.Components.Schemas != nil {
				for k, v := range data.Components.Schemas {
					localDocs.Components.Schemas[k] = v
				}
			}
			if localDocs.Definitions == nil {
				localDocs.Definitions = make(map[string]interface{})
			}
			if data.Definitions != nil {
				for k, v := range data.Definitions {
					localDocs.Definitions[k] = v
				}
			}
			localDocs.Tags = append(localDocs.Tags, data.Tags...)
			mu.Unlock()
		}(url)
	}

	for range services {
		select {
		case <-done:
		case <-time.After(6 * time.Second):
		}
	}

	wg.Wait()

	log.Println("finishing...")

	c.Header("Content-Type", "application/json")
	jsonData, _ := json.Marshal(localDocs)
	c.Header("Content-Length", fmt.Sprintf("%d", len(jsonData)))
	_, err = c.Writer.Write(jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
	}
}

func loadLocalSwagger() (SwaggerDoc, error) {
	data, err := os.ReadFile("./cmd/docs/swagger.json")
	if err != nil {
		return SwaggerDoc{}, err
	}
	var localSwagger SwaggerDoc
	err = json.Unmarshal(data, &localSwagger)
	return localSwagger, err
}

func ScalarHandler(c *gin.Context) {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: "http://localhost:3333/merged-docs",
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}

func RedocHandler(c *gin.Context) {
	html := `
	<!doctype html>
	<html>
	<head>
		<title>ReDoc</title>
		<meta charset="utf-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.css">
	</head>
	<body>
		<redoc spec-url='/merged-docs'></redoc>
		<script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
	</body>
	</html>
	`
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}

func SwaggerRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
}

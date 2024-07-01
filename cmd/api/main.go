package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowAllOrigins: true,
        AllowMethods: []string{"PUT", "PATCH", "DELETE", "POST", "GET", "HEAD"},
        AllowHeaders: []string{"Origin"},
        AllowCredentials: true,
    }))
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H {
            "message": "pong",
        })
    })
    r.Run()
}
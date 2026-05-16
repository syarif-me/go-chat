package main

import (
	"fmt"
	"go-chat/internal/infrastructure/database"
	"go-chat/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitEnv()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test-db", func(c *gin.Context) {
		_, err := database.NewPostgres()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to connect to database",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Connected to database",
		})
	})

	r.Run(":8080")
}

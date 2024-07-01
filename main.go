package main

import (
	"finance/config"
	"finance/db"
	"finance/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var CONFIG config.Config

func main() {
	CONFIG, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return
	}
	db.InitDB(CONFIG.DbName)

	gin.SetMode(CONFIG.GinMode)
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	routes.RegisterRoutes(server)

	server.Run(fmt.Sprintf("%v:%v", CONFIG.Host, CONFIG.Port))

	
}

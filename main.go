package main

import (
	"fmt"
	"go_transaction/config"
	"go_transaction/db"
	"go_transaction/logs"
	"go_transaction/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var CONFIG config.Config

func main() {
	CONFIG, err := config.Load(".envF")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return
	}
	err = logs.Save(CONFIG.LogFilename)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return
	}
	defer logs.CloseLogFile()

	db.InitDB(CONFIG.DbName)
	gin.SetMode(CONFIG.GinMode)
	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer()

	server := gin.New()
	server.Use(logs.CustomGinLogger())
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})
	routes.RegisterRoutes(server)

	log.Printf("server: %v:%v\n", CONFIG.Host, CONFIG.Port)
	server.Run(fmt.Sprintf("%v:%v", CONFIG.Host, CONFIG.Port))

}

package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jimmy/GinAppp/config"
	"github.com/jimmy/GinAppp/routes"
	"github.com/jimmy/GinAppp/app/models"
)

func main() {
	config.LoadEnv()
	config.ConnectDatabase()
	config.InitStorage()

	config.DB.AutoMigrate(&models.User{})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	// Web routes
	routes.RegisterWebRoutes(r)

	// API routes
	routes.RegisterAPIRoutes(r)

	fmt.Printf("🚀 Starting on port %s\n", port)
	r.Run(":" + port)
}

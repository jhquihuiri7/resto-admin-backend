package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"resto-admin-backend/config"
	"resto-admin-backend/routes"
	"time"
)

func main() {

	config.InitFirebase()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Origen permitido
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "Accept", "Content-Length", "X-CSRF-Token", "Token", "session", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes(router)

	// Iniciar el servidor en el puerto 8080
	router.Run(":8080")
}

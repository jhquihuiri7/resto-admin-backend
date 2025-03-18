package main

import (
	"github.com/gin-gonic/gin"

	"resto-admin-backend/config"
	"resto-admin-backend/routes"
)

func main() {

	config.InitFirebase()
	router := gin.Default()

	routes.SetupRoutes(router)

	// Iniciar el servidor en el puerto 8080
	router.Run(":8080")
}

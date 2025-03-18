package routes

import (
	"github.com/gin-gonic/gin"
	"resto-admin-backend/internal/firestore"
	"resto-admin-backend/middlewares"
)

// SetupRoutes configura todas las rutas de la aplicación
func SetupRoutes(router *gin.Engine) {
	router.GET("/showUsers", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"message": "Success"})
	})
	authGroup := router.Group("/auth")
	authGroup.Use(middlewares.FirebaseAuthMiddleware())
	{
		authGroup.GET("/protected", func(ctx *gin.Context) {
			firestore.GetUsers()
			ctx.JSON(200, gin.H{"message": "You have access to this protected route"})
		})
	}
}

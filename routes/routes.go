package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"resto-admin-backend/internal/auth"
	"resto-admin-backend/internal/firestore"
	"resto-admin-backend/internal/structs"
	"resto-admin-backend/middlewares"
	"time"
)

// SetupRoutes configura todas las rutas de la aplicación
func SetupRoutes(router *gin.Engine) {
	router.GET("/showUsers", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"message": "Success"})
	})
	authGroup := router.Group("/auth")
	authGroup.Use(middlewares.FirebaseAuthMiddleware())
	{
		authGroup.GET("/getUser", func(ctx *gin.Context) {
			var userInfo structs.UserInfo
			id := ctx.Query("id")

			data, err := firestore.GetUser(id)
			if err != nil {
				fmt.Println("Error al leer el cuerpo:", err)
				ctx.JSON(400, gin.H{"error": "No se pudo obtener usuario"})
				return
			}
			err = userInfo.FromMap(data)
			if err != nil {
				fmt.Println("Error al parsear data:", err)
				ctx.JSON(400, gin.H{"error": "No se parsear usuario"})
				return
			}
			ctx.JSON(200, userInfo)
		})
		authGroup.GET("/getUsers", func(ctx *gin.Context) {
			users, err := firestore.GetUsers()
			var usersInfo []structs.UserInfo
			if err != nil {
				fmt.Println("Error al leer el cuerpo:", err)
				ctx.JSON(400, gin.H{"error": "No se pudo obtener usuario"})
				return
			}
			for _, user := range users {
				var userInfo structs.UserInfo
				err = userInfo.FromMap(user)
				if err != nil {
					fmt.Println("Error al parsear data:", err)
					ctx.JSON(400, gin.H{"error": "No se parsear usuario"})
					return
				}
				usersInfo = append(usersInfo, userInfo)

			}
			ctx.JSON(200, usersInfo)
		})
		authGroup.POST("/createUser", func(ctx *gin.Context) {
			var user structs.User
			if err := ctx.ShouldBindJSON(&user); err != nil {
				fmt.Println("Error al leer el cuerpo:", err)
				ctx.JSON(400, gin.H{"error": "Formato inválido"})
				return
			}

			auth.CreateUserAuth(user.Email, user.Password)
			ctx.JSON(200, gin.H{"message": "User was created"})
		})
		authGroup.POST("/userInfo", func(ctx *gin.Context) {
			var userInfo structs.UserInfo

			if err := ctx.ShouldBindJSON(&userInfo); err != nil {
				fmt.Println("Error al leer el cuerpo:", err)
				ctx.JSON(400, gin.H{"error": "Formato inválido"})
				return
			}
			userInfo.CreatedDatetime = time.Now()
			userInfo.LastLoginDatetime = time.Now()
			userInfo.SuscriptionExpireDatetime = time.Now().Add(time.Hour * 24 * 15)

			err := userInfo.CreateUserInfo()
			if err != nil {
				fmt.Printf("error creating user info: %v", err)
				ctx.JSON(400, gin.H{"error": "error creating user info"})
				return
			}
			err = userInfo.DeleteId()
			if err != nil {
				fmt.Printf("error deleting ID: %v", err)
				ctx.JSON(400, gin.H{"error": "error deleting ID"})
				return
			}
			//auth.CreateUserAuth(user.Email, user.Password)
			ctx.JSON(200, gin.H{"message": "User Info was created"})
		})
	}
}

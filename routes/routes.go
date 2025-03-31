package routes

import (
	"fmt"
	"resto-admin-backend/internal/auth"
	"resto-admin-backend/internal/firestore"
	"resto-admin-backend/internal/structs"
	"resto-admin-backend/middlewares"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
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
			var userInfo structs.UserInfo

			if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
				fmt.Println("Error al leer el cuerpo para user:", err)
				ctx.JSON(400, gin.H{"error": "Formato inválido crear user auth"})
				return
			}

			_, err := auth.CreateUserAuth(user.Email, user.Password)
			if err != nil {
				fmt.Println("Error al crear usuario:", err)
				ctx.JSON(400, gin.H{"error": "No se pudo crear usuario"})
				return
			}

			if err := ctx.ShouldBindBodyWith(&userInfo, binding.JSON); err != nil {
				fmt.Println("Error al leer el cuerpo para userInfo:", err)
				ctx.JSON(400, gin.H{"error": "Formato inválido crear user info"})
				return
			}
			userInfo.CreatedDatetime = time.Now()
			userInfo.LastLoginDatetime = time.Now()
			userInfo.SuscriptionExpireDatetime = time.Now().Add(time.Hour * 24 * 15)

			err = userInfo.CreateUserInfo()
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

			ctx.JSON(200, gin.H{"message": "User Info was created"})
		})
		authGroup.DELETE("/deleteUser", func(ctx *gin.Context) {
			id := ctx.Query("id")
            user := structs.User{
				Email: id,
			}
			err := user.DeleteUser()
			if err != nil {
				fmt.Printf("error deleting user: %v", err)
				ctx.JSON(400, gin.H{"error": "error deleting user"})
				return
			}

			ctx.JSON(200, gin.H{"message": "User deleted"})
		})
	}
}

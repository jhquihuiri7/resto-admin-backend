package routes

import (
	"fmt"
	"resto-admin-backend/suprabase"
	"strconv"

	"github.com/gin-gonic/gin"
)

func user_routes (userGroup *gin.RouterGroup){
	{
		userGroup.GET("/getRestaurantByUser", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Query("id"))
			if err != nil {
				ctx.JSON(400, map[string]string{"error":err.Error()})
				return
			}
			restaurant, err := suprabase.GetRestaurantByUser(int8(id))
			if err != nil {
				ctx.JSON(400, gin.H{"error":err.Error()})	
				fmt.Println(err)
				return
			}
			ctx.JSON(200, gin.H{"restaurant":restaurant})
		})
		userGroup.GET("/getUser", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Query("id"))
			if err != nil {
				ctx.JSON(400, gin.H{"error":err.Error()})
				return
			}
			fmt.Println("ID", int8(id) )
			user, err := suprabase.GetUser(int8(id))
			if err != nil {
				ctx.JSON(400, gin.H{"error":err.Error()})	
				fmt.Println(err)
				return
			}
			ctx.JSON(200, user)
		})
	}
	 
}
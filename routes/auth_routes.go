package routes

import (
	"github.com/gin-gonic/gin"
)


func auth_routes (authGroup *gin.RouterGroup){
	{
		authGroup.GET("/test", func(ctx *gin.Context) {

		})
	}
	 
}
package routers

import (
	"github.com/gin-gonic/gin"
	"go-mall/controllers"
)


func InitApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("user")
	{
		//BaseRouter.POST("/", controllers.Init)
		BaseRouter.POST("reset-password", controllers.ResetPassword)
	}
	return BaseRouter
}

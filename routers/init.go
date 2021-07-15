package routers

import (
	"github.com/gin-gonic/gin"
	"go-mall/controllers"
)


func InitApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("init")
	{
		BaseRouter.POST("/", controllers.Init)
	}
	return BaseRouter
}

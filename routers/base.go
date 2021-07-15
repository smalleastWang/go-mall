package routers

import (
	"github.com/gin-gonic/gin"
	"go-mall/controllers"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", controllers.Login)
	}
	return BaseRouter
}


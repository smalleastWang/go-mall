package routers

import (
	"github.com/gin-gonic/gin"
	"go-mall/controllers"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("")
	{
		BaseRouter.POST("login", controllers.Login)
		BaseRouter.POST("logout", controllers.Logout)
		BaseRouter.POST("register", controllers.Register)
		BaseRouter.POST("reset-password", controllers.ResetPassword)
	}
	return BaseRouter
}


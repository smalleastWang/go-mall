package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-mall/middleware"
	"go-mall/utils"
)

func Routes() *gin.Engine {
	Router := gin.Default()
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	//gin.Logger("use middleware logger")
	Router.Use(utils.Logger())
	fmt.Println("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	fmt.Println("use middleware cors")

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("register swagger handler")

	// 公共路由
	PublicGroup := Router.Group("")
	PublicGroup.Use(middleware.EnableCookieSession())
	{
		InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	// 鉴权路由
	PrivateGroup := Router.Group("api")
	PrivateGroup.Use(middleware.AuthSession())
	{
		InitApiRouter(PrivateGroup)                   // 注册功能api路由
	}
	fmt.Println("routers register success")

	return Router
}

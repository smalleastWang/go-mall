package routers

import (
	"github.com/gin-gonic/gin"
	"go-mall/controllers"
)
// 地址相关
func AddressApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	// 地址信息增删查改
	BaseRouter := Router.Group("address")
	{
		//BaseRouter.POST("/", controllers.Init)
		BaseRouter.GET("", controllers.GetAddress)
		BaseRouter.POST("", controllers.UpdateAddress)
		BaseRouter.PUT("", controllers.CreateAddress)
		BaseRouter.DELETE("", controllers.DeleteAddress)
	}
	return BaseRouter
}
// 商品相关
func GoodsApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	// 地址信息增删查改
	BaseRouter := Router.Group("goods")
	{
		//BaseRouter.POST("/", controllers.Init)
		BaseRouter.GET(":id", controllers.GetGoodsById)
		BaseRouter.GET("", controllers.GetGoodsList)
		BaseRouter.POST("", controllers.CreateGoods)
		BaseRouter.PUT("", controllers.UpdateGoods)
		BaseRouter.DELETE("", controllers.DeleteGoods)
	}
	return BaseRouter
}
// 订单相关
func OrderApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("order")
	{
		//BaseRouter.POST("/", controllers.Init)
		BaseRouter.GET("", controllers.GetOrder)
		BaseRouter.POST("", controllers.UpdateOrder)
		BaseRouter.PUT("", controllers.CreateOrder)
		BaseRouter.DELETE("", controllers.DeleteOrder)
	}
	return BaseRouter
}


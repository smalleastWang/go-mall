package main

import (
	"go-mall/global"
	"go-mall/routers"
	"log"
)

func main() {

	Router := routers.Routes() // 获取 engine
	//Router.Static("/form-generator", "./resource/page")

	err := Router.Run(":" + global.CONFIG.Http.Port) // 指定端口，运行 Gin
	if err != nil {
		log.Panicln("服务器启动失败：", err.Error())
	}
}

package main

import (
	"go-mall/config"
	"go-mall/global"
	"go-mall/routers"
	"log"
)

func main() {
	// 设置配置文件
	conf, err := config.GetConfig()
	if err != nil {
		log.Panicln("配置文件读取失败：", err.Error())
		return
	}
	global.CONFIG = *conf
	Router := routers.Routes() // 获取 engine
	//Router.Static("/form-generator", "./resource/page")

	err = Router.Run(":" + global.CONFIG.Http.Port) // 指定端口，运行 Gin
	if err != nil {
		log.Panicln("服务器启动失败：", err.Error())
	}
}

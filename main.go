package main

import (
	"ZSTUCA_HardwareRepair/server"
	"ZSTUCA_HardwareRepair/server/conf"
	"github.com/kataras/iris/v12"
)

func main() {

	// iris服务器框架初始化与路由绑定
	app := iris.Default()
	server.Handle(app)

	// 启动服务器
	if err := app.Listen(":" + conf.GetConf().ServerPort); err != nil {
		println("Failed to start server:", err)
	}
}

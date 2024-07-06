package main

import (
	"ZSTUCA_HardwareRepair/server"
	repairModel "ZSTUCA_HardwareRepair/server/repair/model"
	"github.com/kataras/iris/v12"
)

func main() {

	// iris服务器框架初始化与路由绑定
	app := iris.Default()
	server.Handle(app)

	// 启动服务器
	if err := app.Listen(":" + repairModel.GetConf().ServerPort); err != nil {
		println("Failed to start server:", err)
	}
}

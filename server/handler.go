package server

import (
	"ZSTUCA_HardwareRepair/server/database"
	"ZSTUCA_HardwareRepair/server/email"
	repairController "ZSTUCA_HardwareRepair/server/repair/controller"
	repairModel "ZSTUCA_HardwareRepair/server/repair/model"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/robfig/cron/v3"
	"time"
)

// Handle 绑定资源
func Handle(app *iris.Application) {
	// 注册模板引擎
	app.RegisterView(iris.HTML("./webapp/template", ".html"))
	// 静态资源
	handleStatic(app)
	// 错误页面
	handleError(app)
	// 预约报修服务
	handleRepair(app)
}

// handleStatic 绑定静态资源
func handleStatic(app *iris.Application) {
	app.HandleDir("/", "./webapp/static")
	//app.HandleDir("/", "./webapp/static/old") // 兼容旧版
}

// handleError 绑定错误页面
func handleError(app *iris.Application) {
	// 400
	app.OnErrorCode(iris.StatusBadRequest, func(ctx iris.Context) {
		ctx.View("errorPage/400.html")
	})

	// 401
	app.OnErrorCode(iris.StatusUnauthorized, func(ctx iris.Context) {
		ctx.View("errorPage/401.html")
	})

	// 404
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.View("errorPage/404.html")
	})

	// 500
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		// 渲染自定义的 500 错误页面
		ctx.View("errorPage/500.html")
	})

	// 502
	app.OnErrorCode(iris.StatusBadGateway, func(ctx iris.Context) {
		// 渲染自定义的 502 错误页面
		ctx.View("errorPage/502.html")
	})
}

// handleRepair 预约报修服务
func handleRepair(app *iris.Application) {
	// 绑定用户报修申请入口
	mvc.New(app.Party("/")).Handle(new(repairController.ApplyController))

	// 绑定管理员后台
	mvc.New(app.Party("/bs")).Handle(new(repairController.BackstageController))
}

func init() {
	// 创建任务调度器,每天12点调用,提醒滞留请求
	c := cron.New()
	_, err := c.AddFunc("0 12 * * *", func() {
		fmt.Println("reminder for stay request start")
		limit := time.Now().Add(-48 * time.Hour)
		var stayApplies []repairModel.ApplyInfo
		database.Get().Where("create_at < ? AND (admin_id = 0 OR admin_id IS NULL)", limit).Find(&stayApplies)

		// 没有滞留预约信息,终止执行
		if len(stayApplies) == 0 {
			return
		}

		// 向管理员发送提醒接取滞留预约信息的邮件
		emails := repairModel.GetAllAdminsEmail()
		go email.SendInfoEmails(emails, email.ReminderForDelaying, iris.Map{
			"StayApplies": stayApplies,
		})
	})
	if err != nil {
		fmt.Println("cron create err:", err)
	} else {
		fmt.Println("cron successfully create")
	}
	c.Start()
}

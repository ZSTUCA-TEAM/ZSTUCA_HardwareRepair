package repairController

import (
	"ZSTUCA_HardwareRepair/server/conf"
	"ZSTUCA_HardwareRepair/server/database"
	"ZSTUCA_HardwareRepair/server/email"
	repairModel "ZSTUCA_HardwareRepair/server/repair/model"
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)

// ApplyController 报修申请控制器
type ApplyController struct {
}

// PostApply 前端提交报修申请的接口,Post请求方式,相对路径./apply
func (c *ApplyController) PostApply(apply repairModel.ApplyInfo) int {
	fmt.Println("apply get input:", apply)

	// 获取当前时间
	apply.CreateAt = time.Now()

	// 确认是否重复提交
	var prevTimes []time.Time
	database.Get().Model(&repairModel.ApplyInfo{}).Where("name=? AND card_id=?", apply.Name, apply.CardId).Order("create_at DESC").Pluck("create_at", &prevTimes)
	if len(prevTimes) >= 1 && int(apply.CreateAt.Sub(prevTimes[0]).Seconds()) <= 10 {
		fmt.Println("apply repeatedly submitted")
		return iris.StatusConflict
	} // 最近一次提交在10s内，判定为重复提交

	// 存入数据库
	if err := database.Get().Create(&apply).Error; err != nil {
		fmt.Printf("用户请求写入数据库失败:%v\n", err)
		go email.SendInfoEmail(conf.GetConf().Repair.AdminEmail, email.ReminderForInternalError, iris.Map{
			"Text": "用户提交的报修请求写入数据库失败!",
		})
		return iris.StatusInternalServerError
	}
	fmt.Println("apply has written it in the database")

	// 向用户发送收到申请的邮件
	go email.SendInfoEmail(apply.Email, email.MessageForSubmission, iris.Map{
		"ApplyInfo": apply,
	})

	// 利用模板引擎生成内部通知的邮件
	go email.SendInfoEmails(repairModel.GetAllAdminsEmail(), email.ReminderForSubmission, iris.Map{
		"ApplyInfo": apply,
	})

	return iris.StatusCreated
}

package email

import (
	"ZSTUCA_HardwareRepair/server/conf"
	"bytes"
	"fmt"
	"github.com/kataras/iris/v12"
	"html/template"
)

const (
	// MessageForSubmission 报修申请已提交至硬件部后台通知
	MessageForSubmission = iota
	// ReminderForSubmission 有新提交的申请通知
	ReminderForSubmission
	// ReminderForDelaying 接收滞留申请提醒
	ReminderForDelaying
	// MessageForReception 报修申请已被硬件部成员领取通知
	MessageForReception
	// ReminderForReception 已接取预约任务通知
	ReminderForReception
	// MessageForFinish 硬件报修申请已完成通知
	MessageForFinish
	// MessageForAbandoned 报修申请已被放弃通知
	MessageForAbandoned
	// MessageCustomize 自定义内容通知
	MessageCustomize
	// ReminderForInternalError 服务端内部错误通知
	ReminderForInternalError
)

// SendInfoEmails 群发通知邮件
func SendInfoEmails(to []string, emailType int, content iris.Map) {
	// 添加标识通知类型的参数
	content["Type"] = emailType
	content["Title"] = conf.GetConf().Repair.InfoEmailTitle[emailType]

	// 将模板解析到字符串中
	tpl, _ := template.ParseFiles("./webapp/template/repair/InfoEmail.html")
	var body bytes.Buffer
	if err := tpl.Execute(&body, content); err != nil {
		fmt.Println("template for ", conf.GetConf().Repair.InfoEmailForm[emailType], " error:", err)
		return
	}

	// 发送邮件
	if err := SendAll(conf.GetConf().Repair.InfoEmailForm, to, conf.GetConf().Repair.EmailAddr, conf.GetConf().Repair.EmailPort, conf.GetConf().Repair.EmailPassword, conf.GetConf().Repair.InfoEmailTitle[emailType], body.Bytes()); err != nil {
		fmt.Println("email send error:", err)
		return
	}
	fmt.Println("email for ", conf.GetConf().Repair.InfoEmailTitle[emailType], " has sent")
}

// SendInfoEmail 发送通知邮件
func SendInfoEmail(to string, emailType int, content iris.Map) {
	SendInfoEmails([]string{to}, emailType, content)
}

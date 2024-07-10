package repairController

import (
	"ZSTUCA_HardwareRepair/server/conf"
	"ZSTUCA_HardwareRepair/server/database"
	repairModel "ZSTUCA_HardwareRepair/server/repair/model"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// BackstageController 硬件部后端控制器
type BackstageController struct {
}

func (c *BackstageController) BeforeActivation(b mvc.BeforeActivation) {
	b.Router().Use(func(c iris.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.Values().Set("admin", nil)
		} else {
			fmt.Println("管理员jwt:", auth)
			if admin, err := repairModel.GetAdminFromJWT(auth); err != nil {
				fmt.Println("jwt转换管理员对象失败", err)
			} else {
				c.Values().Set("admin", admin)
			}
		}
		c.Next()
	})
}

// GetPub 获取公钥
func (c *BackstageController) GetPub() mvc.Result {
	return mvc.Response{
		Text: conf.GetConf().PublicKeyString,
	}
}

// PostLogin 硬件部后端主界面
func (c *BackstageController) PostLogin(admin repairModel.AdminInfo) mvc.Result {
	fmt.Println("管理员尝试登录,用户名:", admin.Username, " 密码:", admin.Password)

	// 不存在用户名或密码错误
	if !admin.CheckAdminPassword() {
		fmt.Println("管理员用户名或密码错误")
		return mvc.Response{
			ContentType: "application/jwt",
			Code:        iris.StatusUnauthorized,
		}
	}

	// 生成jwt并返回
	jwt, err := admin.GetJWT()
	if err != nil {
		fmt.Println("jwt生成失败,err: ", err)
		return mvc.Response{
			Code: iris.StatusInternalServerError,
		}
	}
	fmt.Println("生成jwt成功")
	return mvc.Response{
		Text: jwt,
	}
}

// GetApplyBy 管理员获取指定页面,Get请求方式
func (c *BackstageController) GetApplyBy(ctx iris.Context, contentType string) mvc.Result {
	// 获取当前管理员对象
	adminInfo := ctx.Values().Get("admin").(*repairModel.AdminInfo)

	if adminInfo == nil {
		fmt.Println("管理员未登录")
		return mvc.Response{
			Code: iris.StatusUnauthorized,
		}
	}

	var applyInfo []repairModel.ApplyInfo
	// 返回相应界面
	if contentType == "new" { // 接取新的委托
		// 读取未接取的预约信息
		database.Get().Where("admin_id IS NULL AND is_abandoned = 0").Find(&applyInfo)
		fmt.Println("管理员", adminInfo.ID, "访问所有新委托")
	} else if contentType == "my" { // 查看已接取的委托
		// 读取当前管理员已接取的预约信息
		database.Get().Where("admin_id = ?", adminInfo.ID).Find(&applyInfo)
		fmt.Println("管理员", adminInfo.ID, "访问自己已接取的委托")
	} else if !adminInfo.IsRootAdmin { // 剩下的查询只有根管理员可以进行
		fmt.Println("管理员", adminInfo.ID, "权限不足")
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	} else if contentType == "all" { // 查看所有委托信息
		fmt.Println("根管理员", adminInfo.ID, "访问所有委托信息")
		database.Get().Preload("Admin").Find(&applyInfo)
	} else {
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}
	return mvc.Response{
		Object: applyInfo,
	}
}

//
//// PostReceive 管理员接取预约,Post请求方式，相对路径./receive
//func (c *BackstageController) PostReceive(ctx iris.Context, applyInfo repairModel.ApplyInfo) mvc.Result {
//	// 获取请求参数中的预约信息ID
//	fmt.Println("admin try to receive the ", applyInfo.ID, " apply")
//
//	// 获取当前管理员对象
//	sess := sessions.Get(ctx)
//	admin, err := tool.CheckAdminPassword(sess.GetString("username"), sess.GetString("password"))
//	if err != nil {
//		fmt.Println("can't get admin info from session")
//		// 当前管理员不存在,需要重新登录
//		return mvc.Response{
//			Code: iris.StatusSeeOther,
//			Path: "/repair/bs",
//		}
//	}
//
//	// 取出当前接取的预约信息
//	if err := database.Get().Where(&applyInfo).Take(&applyInfo).Error; err != nil {
//		fmt.Println("this ID doesn't right")
//		return mvc.Response{
//			Code:        iris.StatusBadRequest,
//			ContentType: "text/html",
//		}
//	}
//
//	if applyInfo.AdminId != 0 {
//		// 该预约已被其他管理员领取
//		fmt.Println("this ", applyInfo.ID, " apply had been received yet")
//		return mvc.Response{
//			Code: iris.StatusOK,
//			Text: "该委托已被接取",
//		}
//	} else if applyInfo.IsAbandoned {
//		// 当前预约已被放弃
//		fmt.Println("this ", applyInfo.ID, " apply had been abandoned")
//		return mvc.Response{
//			Code:        iris.StatusBadRequest,
//			ContentType: "text/html",
//		}
//	}
//
//	// 绑定外键
//	applyInfo.Admin = admin
//	// 改入数据库
//	database.Get().Save(&applyInfo)
//	fmt.Println("admin ", admin.ID, " successfully received the ", applyInfo.ID, " apply")
//
//	// 利用模板引擎生成通知用户申请已被接取的邮件
//	go tool.SendInfoEmail(applyInfo.Email, tool.MessageForReception, iris.Map{
//		"ApplyInfo": applyInfo,
//		"Admin":     admin,
//	})
//
//	// 将详细信息页面发送到该管理员的邮箱中
//	go tool.SendInfoEmail(admin.Email, tool.ReminderForReception, iris.Map{
//		"ApplyInfo": applyInfo,
//		"Admin":     admin,
//	})
//
//	return mvc.Response{
//		Code: iris.StatusOK,
//		Text: "成功接取委托,已将委托人详细个人信息发送至你的邮箱,也可在已接取委托界面查看详细信息.感谢你为硬件部做出的贡献",
//	}
//}
//
//// GetFinishBy 获取提交任务申报界面,Get请求方式,相对路径:./finish/{id: int}
//func (c *BackstageController) GetFinishBy(ctx iris.Context, id int) mvc.Result {
//	// 获取请求参数中的预约信息ID
//	fmt.Println("admin try to finish the ", id, " apply")
//
//	// 获取当前管理员对象
//	sess := sessions.Get(ctx)
//	_, err := tool.CheckAdminPassword(sess.GetString("username"), sess.GetString("password"))
//	if err != nil {
//		fmt.Println("can't get admin info from session")
//		// 当前管理员不存在,需要重新登录
//		return mvc.Response{
//			Code: iris.StatusSeeOther,
//			Path: "/repair/bs",
//		}
//	}
//
//	return mvc.View{
//		Name: "repair/CompleteInfo.html",
//		Data: iris.Map{
//			"ID": id,
//		},
//	}
//}
//
//// PostFinish 管理员提交任务完成申报,Post请求方式,相对路径./finish
//func (c *BackstageController) PostFinish(ctx iris.Context, workList repairModel.WorkList) mvc.Result {
//	// 取出当前登录的管理员对象
//	sess := sessions.Get(ctx)
//	admin, err := tool.CheckAdminPassword(sess.GetString("username"), sess.GetString("password"))
//	if err != nil {
//		fmt.Println("can't get admin info from session")
//		// 当前管理员不存在,需要重新登录
//		return mvc.Response{
//			Code: iris.StatusSeeOther,
//			Path: "/repair/bs",
//		}
//	}
//
//	// 取出当前预约信息对象
//	var applyInfo repairModel.ApplyInfo
//	err = database.Get().Where("id = ?", workList.ApplyId).Take(&applyInfo).Error
//	// 该预约不属于当前管理员或业已完成
//	if err != nil || applyInfo.AdminId != admin.ID || applyInfo.IsFinish || applyInfo.IsAbandoned {
//		fmt.Println("this ", applyInfo.ID, " apply not received by the ", admin.ID, " admin or had been finished or had been abandoned")
//		return mvc.Response{
//			Code:        iris.StatusBadRequest,
//			ContentType: "text/html",
//		}
//	}
//
//	// 创建当前报修任务完成信息
//	workList.Admin = admin
//	workList.Apply = applyInfo
//	database.Get().Create(&workList)
//
//	// 标记当前任务已完成
//	applyInfo.IsFinish = true
//	database.Get().Save(&applyInfo)
//
//	fmt.Println("admin ", admin.ID, " finish the ", applyInfo.ID, " apply")
//
//	// 利用模板引擎生成向用户发送的以便用户核查具体情况的邮件
//	go tool.SendInfoEmail(applyInfo.Email, tool.MessageForFinish, iris.Map{
//		"WorkList": workList,
//	})
//
//	return mvc.Response{
//		Code: iris.StatusOK,
//		Text: "恭喜你成功完成当前任务,再次感谢您作为浙江理工大学计算机协会硬件部的一份子为浙江理工大学做出的贡献!",
//	}
//}
//
//// PostAbandoned 管理员放弃任务,Post请求方式,相对路径./abandoned
//func (c *BackstageController) PostAbandoned(ctx iris.Context, applyInfo repairModel.ApplyInfo) mvc.Result {
//	// 取出当前登录的管理员对象
//	sess := sessions.Get(ctx)
//	admin, err := tool.CheckAdminPassword(sess.GetString("username"), sess.GetString("password"))
//	if err != nil {
//		fmt.Println("can't get admin info from session")
//		// 当前管理员不存在,需要重新登录
//		return mvc.Response{
//			Code: iris.StatusSeeOther,
//			Path: "/repair/bs",
//		}
//	}
//
//	// 取出当前接取的预约信息
//	if err := database.Get().Where(&applyInfo).Take(&applyInfo).Error; err != nil {
//		fmt.Println("this ID doesn't right")
//		return mvc.Response{
//			Code:        iris.StatusBadRequest,
//			ContentType: "text/html",
//		}
//	}
//
//	// 任务信息不正常
//	if admin.ID != applyInfo.AdminId || applyInfo.IsFinish || applyInfo.IsAbandoned {
//		fmt.Println("apply info is wrong")
//		return mvc.Response{
//			Code:        iris.StatusBadRequest,
//			ContentType: "text/html",
//		}
//	}
//
//	// 改入数据库
//	applyInfo.IsAbandoned = true
//	database.Get().Save(&applyInfo)
//	fmt.Println("admin ", admin.ID, " successful abandoned the ", applyInfo.ID, " apply")
//
//	// 向用户发送申请被放弃通知邮件
//	go tool.SendInfoEmail(applyInfo.Email, tool.MessageForAbandoned, iris.Map{
//		"ApplyInfo": applyInfo,
//	})
//
//	return mvc.Response{
//		Code: iris.StatusOK,
//		Text: "成功拒绝委托",
//	}
//}
//
//// PostAdmin 注册管理员,Post请求方式,相对路径./admin
//func (c *BackstageController) PostAdmin(ctx iris.Context, adminInfo repairModel.AdminInfo) mvc.Result {
//	sess := sessions.Get(ctx)
//	rootAdmin, err := tool.CheckAdminPassword(sess.GetString("username"), sess.GetString("password"))
//	if err != nil {
//		fmt.Println("can't get admin info from session")
//		// 当前管理员不存在,需要重新登录
//		return mvc.Response{
//			Code: iris.StatusSeeOther,
//			Path: "/repair/bs",
//		}
//	} else if !rootAdmin.IsRootAdmin {
//		// 当前管理员不具有足够的权限
//		fmt.Println("permission is too low to register admin")
//		return mvc.Response{
//			Code: iris.StatusBadRequest,
//		}
//	}
//
//	fmt.Println("admin ", adminInfo.Username, " try to register")
//
//	// 检验用户名是否重复
//	var tmp repairModel.AdminInfo
//	if rows := database.Get().Where("username = ?", adminInfo.Username).Take(&tmp).RowsAffected; rows != 0 {
//		fmt.Println("duplicate admin username")
//		// 当前用户名重复
//		return mvc.Response{
//			Code: iris.StatusOK,
//			Text: "用户已存在(用户名重复)",
//		}
//	}
//
//	// 加密密码
//	password, _ := bcrypt.GenerateFromPassword([]byte(adminInfo.Password), bcrypt.DefaultCost)
//	adminInfo.Password = string(password)
//
//	// 写入数据库
//	database.Get().Create(&adminInfo)
//
//	fmt.Println("Admin register successfully")
//	return mvc.Response{
//		Code: iris.StatusOK,
//		Text: "用户注册成功",
//	}
//}
//
//// GetSend 管理员向所有未被接取的委托发送自定义信息的通知邮件
//func (c *BackstageController) GetSend(ctx iris.Context, text string) mvc.Result {
//	sess := sessions.Get(ctx)
//	rootAdmin, err := tool.CheckAdminPassword(sess.GetString("username"), sess.GetString("password"))
//	if err != nil {
//		fmt.Println("can't get admin info from session")
//		// 当前管理员不存在,需要重新登录
//		return mvc.Response{
//			Code: iris.StatusSeeOther,
//			Path: "/repair/bs",
//		}
//	} else if !rootAdmin.IsRootAdmin {
//		// 当前管理员不具有足够的权限
//		fmt.Println("permission is too low to register admin")
//		return mvc.Response{
//			Code: iris.StatusBadRequest,
//		}
//	}
//
//	var emails []string
//	database.Get().Model(&repairModel.ApplyInfo{}).Where("admin_id = 0 OR admin_id IS NULL").Pluck("email", &emails) // 获取所有管理员的邮件
//	for i := 0; i < len(emails); i += 10 {
//		tmpEmails := emails[i : i+10]
//		if i+10 >= len(emails) {
//			tmpEmails = emails[i:]
//		}
//
//		// 网易邮箱有并发限制,所以此处不使用协程
//		tool.SendInfoEmails(tmpEmails, tool.MessageCustomize, iris.Map{
//			"Text": text,
//		})
//	}
//
//	return mvc.Response{
//		Code: iris.StatusOK,
//	}
//}

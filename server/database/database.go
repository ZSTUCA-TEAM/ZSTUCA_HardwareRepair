package database

import (
	repairModel "ZSTUCA_HardwareRepair/server/repair/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database 数据库连接单例对象
var database *gorm.DB

// 初始化数据库
func init() {
	if db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: repairModel.GetConf().DataSourceName,
	}), &gorm.Config{}); err != nil {
		fmt.Println("连接数据库失败:", err)
	} else {
		database = db
	}
	fmt.Println("连接数据库成功")

	// 自动迁移
	database.AutoMigrate(repairModel.ApplyInfo{}) // 报修申请信息模型数据表
	database.AutoMigrate(repairModel.AdminInfo{}) // 硬件部成员账号数据表
	database.AutoMigrate(repairModel.WorkList{})  // 报修任务完成信息数据表
	fmt.Println("数据表自动迁移成功")
}

// Get 获取数据库连接单例对象
func Get() *gorm.DB {

	return database
}

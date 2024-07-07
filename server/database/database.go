package database

import (
	"ZSTUCA_HardwareRepair/server/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database 数据库连接单例对象
var database *gorm.DB

// 初始化数据库
func init() {
	if db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: conf.GetConf().DataSourceName,
	}), &gorm.Config{}); err != nil {
		fmt.Println("连接数据库失败:", err)
	} else {
		database = db
	}
	fmt.Println("连接数据库成功")
}

// Get 获取数据库连接单例对象
func Get() *gorm.DB {
	return database
}

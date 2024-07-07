package repairModel

import (
	"ZSTUCA_HardwareRepair/server/conf"
	"ZSTUCA_HardwareRepair/server/database"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// AdminInfo 管理员信息
type AdminInfo struct {
	// ID 序号
	ID uint `json:"id" gorm:"primary_key"`
	// Username 用户名
	Username string `json:"username" gorm:"type:text;not null"`
	// Password 密码(需加密)
	Password string `json:"password" gorm:"type:text;not null"`
	// IsRootAdmin 是否是顶级管理员(硬件部部长)
	IsRootAdmin bool `json:"isRootAdmin" gorm:"not null"`
	// Name 姓名
	Name string `json:"name" gorm:"type:text;not null"`
	// Gender 性别
	Gender string `json:"gender" gorm:"type:text;not null"`
	// Email 邮箱
	Email string `json:"email" gorm:"type:text;not null"`
	// QQ QQ号
	QQ string `json:"qq" gorm:"type:text;not null"`
	// WeChat 微信号
	WeChat string `json:"WeChat" gorm:"type:text;not null"`
	// Phone 手机号
	Phone string `json:"phone" gorm:"type:text;not null"`
}

// CheckAdminPassword 根据管理员用户名和密码取出管理员
func (admin AdminInfo) CheckAdminPassword() bool {
	password := admin.Password
	// 在数据库中查找当前用户名
	fmt.Println("数据库中查询管理员信息")
	rows := database.Get().Where(&admin, "username").Take(&admin).RowsAffected
	if rows == 0 || bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)) != nil {
		// 不存在用户名或密码错误
		return false
	}
	return true
}

func (admin AdminInfo) GetJWT() (signedToken string, err error) {
	claims := jwt.MapClaims{
		"admin": admin,
		"exp":   time.Now().Add(time.Hour * 24 * 7 * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err = token.SignedString(conf.GetConf().PrivateKey)
	return
}

// GetAllAdminsEmail 获取所有管理员的邮箱
func GetAllAdminsEmail() (emails []string) {
	database.Get().Model(&AdminInfo{}).Pluck("email", &emails) // 获取所有管理员的邮件
	return
}

func init() {
	database.Get().AutoMigrate(AdminInfo{})
}

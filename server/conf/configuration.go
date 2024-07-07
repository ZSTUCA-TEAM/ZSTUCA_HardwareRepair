package conf

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
)

// Configuration 配置文件
type Configuration struct {
	ServerPort     string `json:"serverPort"`
	DatabaseName   string `json:"databaseName"`
	DataSourceName string `json:"dataSourceName"`
	Repair         struct {
		EmailAddr      string   `json:"emailAddr"`
		EmailPort      string   `json:"emailPort"`
		EmailPassword  string   `json:"emailPassword"`
		AdminEmail     string   `json:"adminEmail"`
		BsAuthUsername string   `json:"bsAuthUsername"`
		BsAuthPassword string   `json:"bsAuthPassword"`
		InfoEmailForm  string   `json:"infoEmailForm"`
		InfoEmailTitle []string `json:"infoEmailTitle"`
	} `json:"repair"`
	PrivateKey      *rsa.PrivateKey
	PublicKey       *rsa.PublicKey
	PublicKeyString string
}

// conf 配置文件单例对象
var conf Configuration

// GetConf 获取配置文件单例对象
func GetConf() Configuration {
	return conf
}

// init 运行时自动读取配置文件
func init() {
	// 读取配置文件
	fileData, err := os.ReadFile("conf.json")
	if err != nil {
		fmt.Println("无法打开配置文件:", err)
		return
	}

	// 解析配置文件
	if err := json.Unmarshal(fileData, &conf); err != nil {
		fmt.Println("无法解析配置文件:", err)
		return
	}

	// 读取公密钥文件

	if privateKeyBytes, err := os.ReadFile("cert/rsa"); err != nil {
		panic(fmt.Sprintf("无法打开密钥文件:%v", err))
	} else {
		block, _ := pem.Decode(privateKeyBytes)
		if block == nil {
			panic(fmt.Sprintf("无法解析密钥文件:%v", err))
		}
		if conf.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			panic(fmt.Sprintf("无法转换密钥对象:%v", err))
		}
	}
	if publicKeyBytes, err := os.ReadFile("cert/rsa.pub"); err != nil {
		panic(fmt.Sprintf("无法打开公钥文件:%v", err))
	} else {
		block, _ := pem.Decode(publicKeyBytes)
		if block == nil {
			panic(fmt.Sprintf("无法解析公钥文件:%v", err))
		}
		if publicKey, err := x509.ParsePKIXPublicKey(block.Bytes); err != nil {
			panic(fmt.Sprintf("无法转换公钥对象:%v", err))
		} else {
			conf.PublicKey = publicKey.(*rsa.PublicKey)
			conf.PublicKeyString = string(publicKeyBytes)
		}
	}
}

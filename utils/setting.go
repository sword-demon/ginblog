package utils

import "fmt"
import "gopkg.in/ini.v1"

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件错误，请检查文件路径", err)
	}

	// 从配置中加载数据
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	// 如果去配置文件里没有取到值，就默认为debug
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")

}

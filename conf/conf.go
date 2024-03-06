package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
	"web-go/model"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPost     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Print("配置文件读取错误")
	}
	LoadServer(file)
	LoadMysql(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPost, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	model.Database(path)
}

func LoadServer(file *ini.File) {
	service := file.Section("service")
	AppMode = service.Key("AppMode").String()
	HttpPort = service.Key("HttpPort").String()

}

func LoadMysql(file *ini.File) {
	mysql := file.Section("mysql")
	Db = mysql.Key("Db").String()
	DbHost = mysql.Key("DbHost").String()
	DbPost = mysql.Key("DbPost").String()
	DbUser = mysql.Key("DbUser").String()
	DbPassWord = mysql.Key("DbPassWord").String()
	DbName = mysql.Key("DbName").String()
}

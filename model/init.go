package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(conn string) {
	fmt.Print(conn)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		fmt.Print(err)
		panic("数据库连接错误")
	}
	fmt.Print("数据库连接成功")
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)       //表名不加s
	db.DB().SetMaxIdleConns(20)  //设置连接池
	db.DB().SetMaxOpenConns(100) //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}

package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// DB 数据库链接单例
var DB *gorm.DB

func ConnectMysql(connString string) error {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		return fmt.Errorf("数据库连接失败：%s", err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	return nil
}

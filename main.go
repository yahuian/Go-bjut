package main

import (
	"github.com/YahuiAn/Go_bjut/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.Init()
}

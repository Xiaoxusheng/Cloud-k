package db

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func init() {
	db, err := xorm.NewEngine("mysql", "root:admin123@/Cloud-k?charset=utf8")
	if err != nil {
		panic(err)
	}
	log.Println("连接成功")
	Engine = db

}

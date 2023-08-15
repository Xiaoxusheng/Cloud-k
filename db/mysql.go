package db

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func init() {
	//服务器密码为Admin123@
	db, err := xorm.NewEngine("mysql", "root:Admin123@@/cloud-k?charset=utf8")
	if err != nil {
		panic(err)
	}
	log.Println("mysql连接成功")
	Engine = db
	Engine.ShowSQL(true)

}

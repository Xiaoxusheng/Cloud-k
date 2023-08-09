package test

import (
	"Cloud-k/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"xorm.io/xorm"
)

func Test_mysql(t *testing.T) {
	Engine, err := xorm.NewEngine("mysql", "root:admin123@/Cloud-k?charset=utf8")
	if err != nil {
		panic(err)
		return
	}
	log.Println("连接成功")
	exit, err := Engine.IsTableExist("user_basic")
	err = Engine.DropTables("user_basic")
	if err != nil {
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(exit)
	if !exit {
		err := Engine.Sync2(new(models.User_basic))
		if err != nil {
			panic(err)
		}
	}
}

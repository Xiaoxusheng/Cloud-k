package models

import (
	"Cloud-k/db"
	"log"
	"time"
)

type User_basic struct {
	Id        int64     `json:"id"`
	Identity  string    `xorm:"varchar(36) unique" json:"identity" binding:"required "`
	Username  string    `xorm:"varchar(16) unique  'user_name' comment('用户名称')"  json:"username" form:"username"  binding:"required,min=3,max=10" form:"username" `
	Password  string    `xorm:"varchar(36)   notnull  'password' comment('密码')" json:"password" form:"password" binding:"required,min=5,max=10"  form:"password" `
	Email     string    `xorm:"notnull unique 'email' comment('注册邮箱') " json:"email" binding:"required email"`
	CreatedAt time.Time `xorm:"created  notnull 'createdAt' comment('注册时间') " json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated notnull 'updateAt' comment('更新时间')   " json:"updatedAt" `
	DeleteAt  time.Time `xorm:"deleted notnull 'deleteAt'  comment('删除时间') " json:"deleteAt"`
}

func (u User_basic) TableName() string {
	return "user_basic"
}

func GetUser(username, password string) *User_basic {
	user := new(User_basic)
	has, err := db.Engine.Where("username=? and password=?", username, password).Get(user)
	log.Println(has)
	if err != nil {
		panic(err)
	}
	if !has {
		return nil
	}
	return user
}

func GetEmail(email string) bool {
	user := new(User_basic)
	has, err := db.Engine.Where("email", email).Get(user)
	if err != nil {
		panic(err)
	}
	return has
}

func InsertUser(username, password, identity, email string) {
	_, err := db.Engine.Insert(&User_basic{
		Username: username,
		Identity: identity,
		Password: password,
		Email:    email,
	})
	if err != nil {
		panic(err)
	}

}

func GetUserDetail(identity string) *User_basic {
	user := new(User_basic)
	u, err := db.Engine.Where("identity", identity).Get(user)
	if err != nil {
		panic(err)
	}
	if u {
		return user
	}
	return nil

}

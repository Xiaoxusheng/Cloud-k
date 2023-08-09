package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"log"
	"time"
)

type User_basic struct {
	Id        int64     `json:"id"`
	Identity  string    `xorm:"varchar(36) unique" json:"identity" binding:"required "`
	Username  string    `xorm:"varchar(16) unique  'username' comment('用户名称')"  json:"username" form:"username"  binding:"required,min=3,max=10" form:"username" `
	Password  string    `xorm:"varchar(36)   notnull  'password' comment('密码')" json:"password" form:"password" binding:"required,min=5,max=10"  form:"password" `
	Email     string    `xorm:"notnull unique 'email' comment('注册邮箱') " json:"email" binding:"required email"`
	CreatedAt time.Time `xorm:"created  notnull 'createdAt' comment('注册时间') " json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated notnull 'updateAt' comment('更新时间')   " json:"updatedAt" `
	DeleteAt  time.Time `xorm:"deleted notnull 'deleteAt'  comment('删除时间') " json:"deleteAt"`
}

func (u User_basic) TableName() string {
	return "user_basic"
}

func GetUser(username, password string) (*User_basic, error) {
	user := new(User_basic)
	has, err := db.Engine.Where("username=? and password=?", username, password).Get(user)
	log.Println(has, err, username, password)
	if err != nil && !has {
		return nil, err
	}
	return user, nil
}

func GetEmail(email string) bool {
	user := new(User_basic)
	has, err := db.Engine.Where("email", email).Get(user)
	if err != nil {
		panic(uility.ErrorMessage{})
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
		panic(uility.ErrorMessage{})
	}

}

func GetUserDetail(identity string) *User_basic {
	user := new(User_basic)
	u, err := db.Engine.Where("identity", identity).Get(user)
	if err != nil {
		panic(uility.ErrorMessage{})
	}
	if u {
		return user
	}
	return nil

}

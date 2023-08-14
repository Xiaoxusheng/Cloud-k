package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"fmt"
	"log"
	"time"
)

type User_basic struct {
	Id        int       `json:"id"  xorm:"unique"`
	Identity  string    `xorm:"varchar(36) unique" json:"identity,omitempty" binding:"required "`
	Username  string    `xorm:"varchar(16) unique  'username' comment('用户名称')"  json:"username,omitempty" form:"username"  binding:"required,min=3,max=10" form:"username" `
	Password  string    `xorm:"varchar(36)   notnull  'password' comment('密码')" json:"password,omitempty" form:"password" binding:"required,min=5,max=10"  form:"password" `
	Email     string    `xorm:"notnull unique 'email' comment('注册邮箱') " json:"email,omitempty" binding:"required email"`
	CreatedAt time.Time `xorm:"created  notnull 'createdAt' comment('注册时间') " json:"createdAt,omitempty"`
	UpdatedAt time.Time `xorm:"updated notnull 'updateAt' comment('更新时间')   " json:"updatedAt,omitempty" `
	DeleteAt  time.Time `xorm:" notnull 'deleteAt'  comment('删除时间') " json:"deleteAt,omitempty"`
}

func (u User_basic) TableName() string {
	return "user_basic"
}

func GetUser(username, password string) (*User_basic, error) {
	user := new(User_basic)
	has, err := db.Engine.Where("username=? and password=?", username, password).Get(user)
	log.Println(has, err, username, password)
	if err != nil && !has {
		panic(uility.ErrorMessage{
			uility.Error,
			"user_basic表查询出错" + err.Error(),
			"GetUser函数",
			time.Now(),
		})
		return nil, err
	}
	return user, nil
}

func GetEmail(email string) bool {
	user := new(User_basic)
	has, err := db.Engine.Where("email=?", email).Get(user)
	if err != nil && !has {
		panic(uility.ErrorMessage{
			uility.Error,
			"user_basic表查询出错" + err.Error(),
			"GetEmail函数",
			time.Now(),
		})
	}
	return has
}

func GetByUser(username string) bool {
	fmt.Println(username)
	user := new(User_basic)
	has, err := db.Engine.Where("username = ?", username).Get(user)
	if err != nil && !has {
		panic(uility.ErrorMessage{
			uility.Error,
			"user_basic表查询出错" + err.Error(),
			"GetByUser函数",
			time.Now(),
		})
	}
	fmt.Println("username", has)
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
		panic(uility.ErrorMessage{
			uility.Error,
			"user_basic表插入出错" + err.Error(),
			"InsertUser函数",
			time.Now(),
		})
	}

}

func GetUserDetail(identity string) *User_basic {
	user := new(User_basic)
	u, err := db.Engine.Where("identity=?", identity).Get(user)
	if err != nil {
		panic(uility.ErrorMessage{
			uility.Error,
			"user_basic表查询出错" + err.Error(),
			"GetUserDetail函数",
			time.Now(),
		})
	}
	if u {
		return user
	}
	return nil

}

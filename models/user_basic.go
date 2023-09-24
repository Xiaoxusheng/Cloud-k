package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"fmt"
	"log"
	"time"
)

type UserBasic struct {
	Id        int       `json:"id"  xorm:"unique"`
	Identity  string    ` json:"identity,omitempty" binding:"required "`
	Username  string    ` json:"username,omitempty" form:"username"  binding:"required,min=3,max=10" form:"username" `
	Password  string    ` json:"password,omitempty" form:"password" binding:"required,min=5,max=10"  form:"password" `
	Email     string    ` json:"email,omitempty" binding:"required email"`
	CreatedAt time.Time `xorm:"created " json:"created_at,omitempty"`
	UpdatedAt time.Time `xorm:"updated  " json:"updated_att,omitempty" `
	DeleteAt  time.Time `xorm:"deleted  " json:"delete_at,omitempty"`
}

func GetUser(username, password string) (*UserBasic, error) {
	user := new(UserBasic)
	has, err := db.Engine.Where("username=? and password=?", username, password).Get(user)
	log.Println(has, err, username, password)
	if err != nil && !has {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "user_basic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "GetUser函数",
		})
		return nil, err
	}
	return user, nil
}

func GetEmail(email string) bool {
	user := new(UserBasic)
	has, err := db.Engine.Where("email=?", email).Get(user)
	log.Println("email", has)
	if err != nil && !has {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "user_basic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "GetEmail函数",
		})
	}

	return has
}

func GetByUser(username string) bool {
	fmt.Println(username)
	user := new(UserBasic)
	has, err := db.Engine.Where("username = ?", username).Get(user)
	if err != nil && !has {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "user_basic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "GetByUser函数",
		})
	}
	fmt.Println("username", has)
	return has
}

func InsertUser(username, password, identity, email string) {
	_, err := db.Engine.Insert(&UserBasic{
		Username: username,
		Identity: identity,
		Password: password,
		Email:    email,
	})
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "user_basic表插入出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "InsertUser函数",
		})
	}

}

func GetUserDetail(identity string) *UserBasic {
	user := new(UserBasic)
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

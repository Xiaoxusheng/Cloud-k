package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"time"
)

type CasbinBasic struct {
	CasbinIdentity string    `json:"casbin_identity"`
	UserIdentity   string    `json:"user_identity"`
	Id             int64     `json:"id"`
	CreatedAt      time.Time `json:"created_at" xorm:"created"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"updated"`
	DeletedAt      time.Time `json:"deleted_at " `
}

func GetUserPermission(identity string) {
	c := new(CasbinBasic)
	ok, err := db.Engine.Where("user_identity=?", identity).Get(c)
	if err != nil || !ok {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "CasbinBasic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "GetUserPermission函数",
		})
	}
}

// 用户
func InsertUserPermission(identity string) {
	_, err := db.Engine.Insert(CasbinBasic{
		CasbinIdentity: uility.User,
		UserIdentity:   identity,
	})
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "CasbinBasic表插入出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "InsertUserPermission函数",
		})
	}
}

func InsertAdmin(identity string) {
	_, err := db.Engine.Insert(CasbinBasic{
		CasbinIdentity: uility.Admin,
		UserIdentity:   identity,
	})
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "CasbinBasic表插入出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "InsertAdmin函数",
		})
	}
}

func UpdatePermission(id, CasbinIdentity string) {
	_, err := db.Engine.Where("identity=?", id).Update(&CasbinBasic{
		CasbinIdentity: CasbinIdentity,
	})
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "CasbinBasic表插入出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "UpdatePermissionn函数",
		})
	}
}

package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"time"
)

type User_repository struct {
	Id                 int       `json:"id"`
	Identity           string    `json:"identity"`
	UserIdentity       string    `json:"user_identity"`       //用户唯一标识
	ParentId           int       `json:"parent_id"`           //父级id
	RepositoryIdentity string    `json:"repository_identity"` //
	Ext                int       `json:"ext"`                 //文件夹或者文件类型
	Name               string    `json:"name"`
	Size               int       `json:"size"`
	CreatedAt          time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt          time.Time `json:"updatedAt" xorm:"updated"`
	DeletedAt          time.Time `json:"deletedAt" xorm:"deleted"`
}

func GetByUserRepository(user_identity, repository_Identity, name string, parent_id int, ext int) bool {
	userRepository := new(User_repository)
	has, err := db.Engine.Where("user_identity=?,parent_id=?,repository_identity=?,ext=?,name=?", user_identity, parent_id, repository_Identity, ext, name).Get(userRepository)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "GetByUserRepository函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "查询User_repository" + err.Error(),
		})
	}
	return has

}

func InsertUserRepository(user *uility.UserRepositorySave) {
	_, err := db.Engine.Insert(&User_repository{
		Identity:           uility.GetUuid(),
		UserIdentity:       user.UserIdentity,
		ParentId:           user.ParentId,
		RepositoryIdentity: user.RepositoryIdentity,
		Ext:                user.Ext,
		Name:               user.Name,
	})

	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "InsertUserRepository函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "User_repository表插入" + err.Error(),
		})
	}

}

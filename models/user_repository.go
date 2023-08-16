package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"log"
	"time"
)

type User_repository struct {
	Id                 int       `json:"id"`
	Identity           string    `json:"identity"`
	UserIdentity       string    `json:"user_identity"`       //用户唯一标识
	ParentId           int       `json:"parent_id"`           //父级id
	RepositoryIdentity string    `json:"repository_identity"` //
	Ext                string    `json:"ext"`                 //文件夹或者文件类型
	Name               string    `json:"name"`
	CreatedAt          time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt          time.Time `json:"updatedAt" xorm:"updated"`
	DeletedAt          time.Time `json:"deletedAt" xorm:"deleted"`
}

func GetByUserRepository(user_identity, repository_Identity string) bool {
	userRepository := new(User_repository)
	has, err := db.Engine.Where("user_identity=? and repository_identity=?", user_identity, repository_Identity).Get(userRepository)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "GetByUserRepository函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "查询User_repository" + err.Error(),
		})
	}
	log.Println(has)
	return has

}

func InsertUserRepository(user *uility.UserRepositorySave) {
	_, err := db.Engine.Insert(&User_repository{
		Identity:           uility.GetUuid(),
		UserIdentity:       user.UserIdentity,
		ParentId:           user.Parent_id,
		RepositoryIdentity: user.Repository_identity,
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

func GetFileList(page, number int, user_identity, parent_id string) []uility.UserRepositoryFileList {
	UserRepositoryFileList := make([]uility.UserRepositoryFileList, 0)
	err := db.Engine.Table("user_repository").Join("left", "repository_pool", "user_repository.repository_identity=repository_pool.identity and user_repository.user_identity=? and user_repository.parent_id=?", user_identity, parent_id).Limit(number, (page-1)*number).Find(&UserRepositoryFileList)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "GetFileList函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "User_repository表查询" + err.Error(),
		})
	}
	return UserRepositoryFileList
}

func GetByName(name string) bool {
	User_repository := new(User_repository)
	has, err := db.Engine.Where("name=?", name).Get(User_repository)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "GetByName函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "User_repository表查询" + err.Error(),
		})
	}
	return has
}

func GetByIdentity(identity, userIdentity string) bool {
	user_repository := new(User_repository)
	has, err := db.Engine.Where("identity=? and parent_id=(select parent_id from user_repository where user_identity=?)", identity, userIdentity).Get(user_repository)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "GetByName函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "User_repository表查询" + err.Error(),
		})
	}
	return has
}

func UpDateFileName(name, identity, userIdentity string) {
	_, err := db.Engine.Where("user_identity=? and identity=?", userIdentity, identity).Update(&User_repository{Name: name})
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "UpDateFileName函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "User_repository表更新" + err.Error(),
		})
	}
}

func GetByNameParentId(name, parent_id string) bool {
	user_repository := new(User_repository)
	has, err := db.Engine.Where("name=? and parent_id=?", name, parent_id).Get(user_repository)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "GetByNameParentId函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "User_repository表查询" + err.Error(),
		})
	}
	return has
}

func InsertFolder(user_identity, name string, parent_id int) {
	_, err := db.Engine.Insert(&User_repository{
		Identity:     uility.GetUuid(),
		UserIdentity: user_identity,
		ParentId:     parent_id,
		Ext:          "",
		Name:         name,
	})
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "InsertFolder函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "User_repository表插入" + err.Error(),
		})
	}
}

func DeleteFile(identity, user_identity string) {
	user_repository := new(User_repository)
	_, err := db.Engine.Where("identity=? and user_identity=? ", identity, user_identity).Delete(user_repository)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "DeleteFile函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "User_repository表删除" + err.Error(),
		})
	}
}

func GetByIdentityUserIdentity(identity, user_identity string) bool {
	user_repository := new(User_repository)
	has, err := db.Engine.Where("identity=? and user_identity=? ", identity, user_identity).Get(user_repository)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDetails:     "DeleteFile函数",
			ErrorTime:        time.Now(),
			ErrorDescription: "User_repository表删除" + err.Error(),
		})
	}
	return has
}

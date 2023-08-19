package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"time"
)

type ShareBasic struct {
	Id                     int       `json:"id"`
	Identity               string    `json:"identity"`
	UserIdentity           string    `json:"user_identity"` //用户唯一标识
	UserRepositoryIdentity string    `json:"user_repository_identity"`
	RepositoryIdentity     string    `json:"repository_identity"` //文件唯一表标识
	ExpiredTime            int       `json:"expired_time"`        //分享失效时间
	ClickNum               int       `json:"click_num"`           //分享次数
	CreatedAt              time.Time `json:"created_at" xorm:"created"`
	UpdatedAt              time.Time `json:"updated_at" xorm:"updated"`
	DeletedAt              time.Time `json:"deleted_at" xorm:"deleted"`
}

func InsertShareBasic(identity, user_identity, user_repository_identity, repository_identity string, expired_time int) {
	_, err := db.Engine.Insert(&ShareBasic{
		Identity:               identity,
		UserIdentity:           user_identity,
		UserRepositoryIdentity: user_repository_identity,
		RepositoryIdentity:     repository_identity,
		ExpiredTime:            expired_time,
		ClickNum:               0,
	})
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "share_basic表插入出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "InsertShareBasic函数",
		})
	}
}

func GetIdentity(identity string) bool {
	ShareBasic := new(ShareBasic)
	has, err := db.Engine.Where("identity=?", identity).Get(ShareBasic)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "share_basic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "GetIdentity函数",
		})
	}
	return has
}

func UpdateClickNum(identity string) {
	//Where("identity=?", identity).Update("click_num=click_num+1")
	_, err := db.Engine.Exec("update share_basic set click_num=click_num+1 where identity=? ", identity)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "share_basic表更新出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "UpdateClickNum函数",
		})
	}
}

func GetShareBasicDetail(identity string) (bool, *uility.ShareBasicFileDetail) {
	ShareBasicFileDetail := new(uility.ShareBasicFileDetail)
	has, err := db.Engine.Table("share_basic").Select("share_basic.repository_identity,repository_pool.name,repository_pool.size,repository_pool    .ext,repository_pool.path").
		Join("left", "user_basic", "user_basic.identity=share_basic.user_identity").
		Join("left", "repository_pool", "repository_pool.identity=share_basic.repository_identity").Where("share_basic.identity= ?", identity).Get(ShareBasicFileDetail)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "share_basic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "GetShareBasicDetail函数",
		})
	}
	return has, ShareBasicFileDetail
}

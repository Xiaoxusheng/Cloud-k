package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"time"
)

type Repository_pool struct {
	Id         int       `json:"id"`
	Identity   string    `json:"identity"`
	Hash       string    `json:"hash"` //内容唯一hash
	Name       string    `json:"name"`
	Ext        string    `json:"ext"`
	Size       int64     `json:"size"`
	Path       string    `json:"path"`
	Created_at time.Time `json:"created_At" xorm:"created"`
	Updated_at time.Time `json:"updated_At" xorm:"updated"`
	Deleted_at time.Time `json:"deleted_At " `
}

func GetByHash(hash string) (bool, error) {
	repository_pool := new(Repository_pool)
	ext, err := db.Engine.Where("hash=?", hash).Get(repository_pool)
	if err != nil {
		return ext, err
	}
	return ext, nil

}

func InsertFile(hash, name, ext, path string, size int64) {
	_, err := db.Engine.Insert(
		&Repository_pool{
			Identity: uility.GetUuid(),
			Hash:     hash,
			Name:     name,
			Ext:      ext,
			Size:     size,
			Path:     "https://cloud-k-1308109276.cos.ap-nanjing.myqcloud.com/" + path,
		})
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "user_basic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "InsertFile函数",
		})
	}
}

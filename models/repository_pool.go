package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"time"
)

type RepositoryPool struct {
	Id        int       `json:"id"`
	Identity  string    `json:"identity"`
	Hash      string    `json:"hash"` //内容唯一hash
	Name      string    `json:"name"`
	Ext       string    `json:"ext"`
	Size      int64     `json:"size"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
	DeletedAt time.Time `json:"deleted_at " `
}

func GetByHash(hash string) (bool, error) {
	repository_pool := new(RepositoryPool)
	ext, err := db.Engine.Where("hash=?", hash).Get(repository_pool)
	if err != nil {
		return ext, err
	}
	return ext, nil

}

func InsertFile(hash, name, ext, path string, size int64) {
	_, err := db.Engine.Insert(
		&RepositoryPool{
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
			ErrorDescription: "repository_pool表插入出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "InsertFile函数",
		})
	}
}

func GetByRepositoryPool(identity string) (bool, *RepositoryPool) {
	repository_pool := new(RepositoryPool)
	has, err := db.Engine.Where("identity=?", identity).Get(repository_pool)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "repository_pool表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "GetByRepositoryPool函数",
		})
	}
	return has, repository_pool
}

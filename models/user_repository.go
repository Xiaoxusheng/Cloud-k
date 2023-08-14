package models

import (
	"time"
)

type User_repository struct {
	Id                  int       `json:"id"`
	Identity            string    `json:"identity"`
	User_Identity       string    `json:"user_Identity"`       //用户唯一标识
	Parent_id           int       `json:"parent_id"`           //父级id
	Repository_identity string    `json:"repository_Identity"` //
	Ext                 int       `json:"ext"`                 //文件夹或者文件类型
	Name                string    `json:"name"`
	Size                int       `json:"size"`
	Created_at          time.Time `json:"created_At" xorm:"created"`
	Updated_at          time.Time `json:"updated_At" xorm:"updated"`
	Deleted_at          time.Time `json:"deleted_At"`
}

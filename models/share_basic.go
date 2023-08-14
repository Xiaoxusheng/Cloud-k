package models

import "time"

type Share_basic struct {
	Id                  int       `json:"id"`
	Identity            string    `json:"identity"`
	User_Identity       string    `json:"user_identity"`       //用户唯一标识
	Repository_identity string    `json:"repository_identity"` //文件唯一表标识
	Expired_time        time.Time `json:"expired_time"`        //分享失效时间
	Click_num           int       `json:"click_num"`           //分享次数
	Created_at          time.Time `json:"created_At" xorm:"created"`
	Updated_at          time.Time `json:"updated_At" xorm:"updated"`
	Deleted_at          time.Time `json:"deleted_At"`
}

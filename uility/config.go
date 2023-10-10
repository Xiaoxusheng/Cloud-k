package uility

import (
	"time"
)

var (
	//token密匙
	MySigningKey = []byte("welcome to use Cloud-kAuth:Mr.Lei")
	//腾讯云cos
	SECRETID  = ""
	SECRETKEY = ""
	key       = "cloud-k"
	//权限唯一id
	User  = "93bb79d2-5045-5cd9-acd7-82980dea0e10"
	Admin = "f07cf079-3139-59ea-9b21-0ffbfbbab3ed"
	Root  = "9525b839-5074-5a50-999f-6d118c1f4857"
)

var Count = 0

type ErrorMessage struct {
	ErrorType        string    `json:"errorType,omitempty"`        //错误类型
	ErrorDescription string    `json:"errorDescription,omitempty"` //细节描述
	ErrorDetails     string    `json:"errorDetails,omitempty"`     // 错误详情
	ErrorTime        time.Time `json:"errorTime,omitempty"`        //时间
}

// 错误级别
const (
	Info      = "100"
	Warning   = "300"
	Error     = "400"
	Critical  = "500"
	Emergency = "999"
)

type UserRepositorySave struct {
	UserIdentity       string `json:"uer_identity,omitempty"  form:"user_identity"`
	Identity           string `json:"identity,omitempty" form:"identity"`
	ParentId           int    `json:"parent_id" binding:"-" form:"parent_id"`
	RepositoryIdentity string `json:"repository_identity" binding:"required" form:"repository_identity"`
	Ext                string `json:"ext" binding:"required" form:"ext"`
	Name               string `json:"name" binding:"required" form:"name"`
	Size               int    `json:"size" form:"size"`
}

type UserRepositoryFileList struct {
	Id                 int    `json:"id"`
	Hash               string `json:"hash"` //内容唯一hash
	Path               string `json:"path"`
	UserIdentity       string `json:"uer_identity,omitempty"  form:"user_identity"`
	Identity           string `json:"identity,omitempty" form:"identity"`
	ParentId           int    `json:"parent_id" binding:"-" form:"parent_id"`
	RepositoryIdentity string `json:"repository_identity" binding:"required" form:"repository_identity"`
	Ext                string `json:"ext" binding:"required" form:"ext"`
	Name               string `json:"name" binding:"required" form:"name"`
	Size               int    `json:"size" form:"size"`
}

type ShareBasicFileDetail struct {
	Name               string `json:"name,"`
	Size               int    `json:"size "`
	Path               string `json:"path"`
	Ext                string `json:"ext"`
	RepositoryIdentity string `json:"repository_identity"`
}

type Userinfo struct {
	Id             int       `json:"id"  xorm:"unique"`
	Identity       string    ` json:"identity,omitempty" binding:"required "`
	Username       string    ` json:"username,omitempty" form:"username"  binding:"required,min=3,max=10" form:"username" `
	Password       string    ` json:"password,omitempty" form:"password" binding:"required,min=5,max=10"  form:"password" `
	Email          string    ` json:"email,omitempty" binding:"required email"`
	CreatedAt      time.Time `xorm:"created " json:"created_at,omitempty"`
	UpdatedAt      time.Time `xorm:"updated  " json:"updated_att,omitempty" `
	DeleteAt       time.Time `xorm:"deleted  " json:"delete_at,omitempty"`
	CasbinIdentity string    `json:"casbinIdentity"`
}

//func i() {
//	var k *int
//	fmt.Println(k)
//	m := 9
//	k = &m
//}

package uility

import (
	"fmt"
	"time"
)

var MySigningKey = []byte("welcome to use Cloud-kAuth:Mr.Lei")

var SECRETID = "AKIDfBRQAdpkPnukJceOr52JK4XjeIgmb9RS"
var SECRETKEY = "lFRTFkzziAMIyNULvEG0VkofGahZBWaN"

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
	UserIdentity        string `json:"uer_identity,omitempty"  form:"user_identity"`
	Identity            string `json:"identity,omitempty" form:"identity"`
	Parent_id           int    `json:"parent_id" binding:"-" form:"parent_id"`
	Repository_identity string `json:"repository_identity" binding:"required" form:"repository_identity"`
	Ext                 string `json:"ext" binding:"required" form:"ext"`
	Name                string `json:"name" binding:"required" form:"name"`
	Size                int    `json:"size" form:"size"`
}

type UserRepositoryFileList struct {
	Id                  int    `json:"id"`
	Hash                string `json:"hash"` //内容唯一hash
	Path                string `json:"path"`
	UserIdentity        string `json:"uer_identity,omitempty"  form:"user_identity"`
	Identity            string `json:"identity,omitempty" form:"identity"`
	Parent_id           int    `json:"parent_id" binding:"-" form:"parent_id"`
	Repository_identity string `json:"repository_identity" binding:"required" form:"repository_identity"`
	Ext                 string `json:"ext" binding:"required" form:"ext"`
	Name                string `json:"name" binding:"required" form:"name"`
	Size                int    `json:"size" form:"size"`
}

type ShareBasicFileDetail struct {
	Name               string `json:"name,"`
	Size               int    `json:"size "`
	Path               string `json:"path"`
	Ext                string `json:"ext"`
	RepositoryIdentity string `json:"repository_identity"`
}

func i() {
	var k *int
	fmt.Println(k)
	m := 9
	k = &m
}

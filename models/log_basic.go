package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"fmt"
	"time"
)

type LogBasic struct {
	Id            int           `json:"id,omitempty"`
	Ip            string        `json:"ip,omitempty"`
	StatusCode    int           `json:"status_Code,omitempty"`
	TimeConsuming time.Duration `json:"time_Consuming,omitempty"`
	UserIdentity  string        `json:"userIdentity,omitempty"`
	Identity      string        `json:"identity,omitempty"`
	Methods       string        `json:"methods,omitempty"`
	Path          string        `json:"path,omitempty"`
	RequestTime   time.Time     `json:"requestTime"`
	Role          string        `json:"role,omitempty"`
}

func GetLogBasicList() []LogBasic {
	list := make([]LogBasic, 0)
	err := db.Engine.Table("log_basic").Find(&list)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "CapacityBasic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "GetLogBasicList函数",
		})
	}
	return list
}

func InsertLog(l *LogBasic) {
	_, err := db.Engine.Insert(l)
	if err != nil {
		fmt.Println(err)
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "CapacityBasic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "InsertLog函数",
		})
	}
}
func GetUserLog() []LogBasic {
	list := make([]LogBasic, 0)
	err := db.Engine.Where("role=?", uility.User).Find(&list)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "CapacityBasic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "GetUserLog函数",
		})
	}
	return list
}

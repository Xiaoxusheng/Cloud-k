package models

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"time"
)

type CapacityBasic struct {
	TotalCapacity    int       `json:"total_capacity,omitempty"`
	ResidualCapacity int       `json:"residual_capacity,omitempty"`
	Id               int       `json:"id,omitempty"`
	Recharge         bool      `json:"recharge,omitempty"`
	Identity         string    `json:"identity,omitempty"`
	UserIdentity     string    `json:"user_identity"`
	CreatedAt        time.Time `json:"created_at" xorm:"created"`
	UpdatedAt        time.Time `json:"updated_at" xorm:"updated"`
	DeletedAt        time.Time `json:"deleted_at " `
}

func InsertCapacityBasic(c []*CapacityBasic) {
	_, err := db.Engine.Insert(c)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "CapacityBasic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "InsertCapacityBasic函数",
		})
	}
}

func GetResidualCapacityList() []CapacityBasic {
	list := make([]CapacityBasic, 0)
	err := db.Engine.Find(&list)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorDescription: "CapacityBasic表查询出错" + err.Error(),
			ErrorTime:        time.Now(),
			ErrorDetails:     "GetResidualCapacityList函数",
		})
	}
	return list
}

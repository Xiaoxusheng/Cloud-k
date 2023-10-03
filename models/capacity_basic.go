package models

import "time"

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

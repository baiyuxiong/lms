package models

import (
	"time"
)

type LeaveTransfer struct {
	Id        int64     `json:"id" xorm:"BIGINT(20)"`
	LeaveId   int       `json:"leave_id" xorm:"not null INT(11)"`
	AssignFr  int       `json:"assign_fr" xorm:"not null INT(11)"`
	AssignTo  int       `json:"assign_to" xorm:"not null INT(11)"`
	IsAgree   int       `json:"is_agree" xorm:"not null TINYINT(4)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
}

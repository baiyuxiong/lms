package models

import (
	"time"
)

type TaskTransfer struct {
	Id        int64     `json:"id" xorm:"BIGINT(20)"`
	TaskId    int       `json:"task_id" xorm:"not null INT(11)"`
	AssignFr  int       `json:"assign_fr" xorm:"not null INT(11)"`
	AssignTo  int       `json:"assign_to" xorm:"not null INT(11)"`
	IsRead    int       `json:"is_read" xorm:"not null TINYINT(4)"`
	Info      string    `json:"info" xorm:"not null TEXT"`
	Progress  int       `json:"progress" xorm:"not null TINYINT(4)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
}

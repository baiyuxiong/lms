package models

import (
	"time"
)

type Leave struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId         int       `json:"user_id" xorm:"not null INT(11)"`
	InChargeUserId int       `json:"in_charge_user_id" xorm:"not null INT(11)"`
	FlowId         string    `json:"flow_id" xorm:"not null VARCHAR(8)"`
	StepId         string    `json:"step_id" xorm:"not null VARCHAR(8)"`
	Status         int       `json:"status" xorm:"not null TINYINT(4)"`
	LeaveType      int       `json:"leave_type" xorm:"not null TINYINT(4)"`
	Reason         string    `json:"reason" xorm:"not null VARCHAR(256)"`
	Address        string    `json:"address" xorm:"not null VARCHAR(128)"`
	OnwayDate      int       `json:"onway_date" xorm:"not null INT(11)"`
	Filepath       string    `json:"filepath" xorm:"not null VARCHAR(256)"`
	Startdate      time.Time `json:"startdate" xorm:"not null DATETIME"`
	Enddate        time.Time `json:"enddate" xorm:"not null DATETIME"`
	Duration       string    `json:"duration" xorm:"not null DECIMAL(10,1)"`
	CreatedAt      time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"not null DATETIME"`
}

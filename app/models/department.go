package models

import (
	"time"
)

type Department struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ParentId  int       `json:"parent_id" xorm:"not null INT(11)"`
	Depth     int       `json:"depth" xorm:"not null INT(11)"`
	LeaderId  int       `json:"leader_id" xorm:"not null INT(11)"`
	Name      string    `json:"name" xorm:"not null VARCHAR(256)"`
	Info      string    `json:"info" xorm:"not null TEXT"`
	IsDelete  int       `json:"is_delete" xorm:"not null default 0 TINYINT(4)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
}

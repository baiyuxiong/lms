package models

import (
	"time"
)

type Project struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	CompanyId int       `json:"company_id" xorm:"not null INT(11)"`
	OwnerId   int       `json:"owner_id" xorm:"not null INT(11)"`
	Name      string    `json:"name" xorm:"not null VARCHAR(256)"`
	Info      string    `json:"info" xorm:"not null TEXT"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
}

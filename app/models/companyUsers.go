package models

import (
	"time"
)

type CompanyUsers struct {
	CompanyId int       `json:"company_id" xorm:"not null pk unique(company_id) INT(11)"`
	UserId    int       `json:"user_id" xorm:"not null pk unique(company_id) INT(11)"`
	Status    int       `json:"status" xorm:"not null TINYINT(4)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
}

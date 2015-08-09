package models

import (
	"time"
)

type Company struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OwnerId          int       `json:"owner_id" xorm:"not null INT(11)"`
	UserCheckCount   int       `json:"user_check_count" xorm:"not null INT(11)"`
	UserUncheckCount int       `json:"user_uncheck_count" xorm:"not null INT(11)"`
	ProjectCount     int       `json:"project_count" xorm:"not null INT(11)"`
	Name             string    `json:"name" xorm:"not null VARCHAR(128)"`
	Info             string    `json:"info" xorm:"not null TEXT"`
	Logo             string    `json:"logo" xorm:"not null VARCHAR(256)"`
	Phone            string    `json:"phone" xorm:"VARCHAR(12)"`
	Address          string    `json:"address" xorm:"VARCHAR(256)"`
	CreatedAt        time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt        time.Time `json:"updated_at" xorm:"not null DATETIME"`
}

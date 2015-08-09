package models

import (
	"time"
)

type SmsCode struct {
	Username  string    `json:"username" xorm:"not null pk unique CHAR(11)"`
	Code      string    `json:"code" xorm:"not null CHAR(6)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
}

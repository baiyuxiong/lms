package models

import (
	"time"
)

type Complain struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId       int       `json:"user_id" xorm:"not null INT(11)"`
	ComplainId   int       `json:"complain_id" xorm:"not null INT(11)"`
	ComplainType int       `json:"complain_type" xorm:"not null INT(11)"`
	Info         string    `json:"info" xorm:"not null TEXT"`
	CreatedAt    time.Time `json:"created_at" xorm:"not null DATETIME"`
	Updatedat    time.Time `json:"UpdatedAt" xorm:"not null DATETIME"`
}

package models

import (
	"time"
)

type News struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	CompanyId int       `json:"company_id" xorm:"not null INT(11)"`
	ProjectId int       `json:"project_id" xorm:"not null INT(11)"`
	OwnerId   int       `json:"owner_id" xorm:"not null INT(11)"`
	Title     string    `json:"title" xorm:"not null VARCHAR(255)"`
	Content   string    `json:"content" xorm:"not null TEXT"`
	CreatedAt time.Time `json:"created_at" xorm:"not null default '0000-00-00 00:00:00' TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default '0000-00-00 00:00:00' TIMESTAMP"`
}

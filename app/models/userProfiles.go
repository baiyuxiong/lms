package models

import (
	"time"
)

type UserProfiles struct {
	UserId         int       `json:"user_id" xorm:"not null pk unique index INT(11)"`
	DepartmentId   int       `json:"department_id" xorm:"not null INT(11)"`
	Gender         int       `json:"gender" xorm:"not null TINYINT(4)"`
	Level          int       `json:"level" xorm:"not null INT(11)"`
	EmploymentType int       `json:"employment_type" xorm:"not null TINYINT(4)"`
	Name           string    `json:"name" xorm:"not null VARCHAR(32)"`
	Phone          string    `json:"phone" xorm:"not null VARCHAR(16)"`
	Phone1         string    `json:"phone1" xorm:"not null VARCHAR(16)"`
	Position       string    `json:"position" xorm:"not null VARCHAR(64)"`
	JoinDate       time.Time `json:"join_date" xorm:"not null DATE"`
	Email          string    `json:"email" xorm:"not null VARCHAR(64)"`
	BirthDate      time.Time `json:"birth_date" xorm:"not null DATE"`
	CreatedAt      time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"not null DATETIME"`
}

package models

type Level struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name     string `json:"name" xorm:"not null VARCHAR(32)"`
	Level    int    `json:"level" xorm:"not null INT(11)"`
	IsDelete int    `json:"is_delete" xorm:"not null default 0 TINYINT(4)"`
}

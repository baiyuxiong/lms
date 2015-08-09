package models

type Flow struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	InCharge    string `json:"in_charge" xorm:"not null VARCHAR(64)"`
	Type        int    `json:"type" xorm:"not null TINYINT(4)"`
	Step        int    `json:"step" xorm:"not null INT(11)"`
	IfCondition string `json:"if_condition" xorm:"not null VARCHAR(128)"`
	NextStep    int    `json:"next_step" xorm:"not null INT(11)"`
	Info        string `json:"info" xorm:"not null VARCHAR(64)"`
}

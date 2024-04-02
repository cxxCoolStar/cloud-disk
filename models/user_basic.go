package models

type UserBasic struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	Identity string `xorm:"not null comment('用户身份') VARCHAR(255)"`
	Name     string
	Password string
	Email    string
}

func (table UserBasic) TableName() string {
	return "user_basic"
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"varchar(32);not null;comment:'用户id'"`
	Nickname string `gorm:"type:varchar(20);not null;comment:'昵称'"`
	Password string `gorm:"type:varchar(100);not null;comment:'密码'"`
	Phone    string `gorm:"type:varchar(20);comment:'手机号'"`
	WeChat   string `gorm:"type:varchar(20);comment:'微信'"`
	Email    string `gorm:"type:varchar(100);comment:'邮箱'"`
}

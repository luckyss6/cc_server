package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserID   string `gorm:"type:varchar(64);uniqueIndex;not null"`
	Username string `gorm:"type:varchar(20);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(64);not null"`
	Email    string `gorm:"type:varchar(40);uniqueIndex;not null"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}

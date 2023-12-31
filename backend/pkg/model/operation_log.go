package model

import "gorm.io/gorm"

type OperationLog struct {
	gorm.Model
	UserID    string `gorm:"type:varchar(64);not null"`
	Username  string `gorm:"type:varchar(20);not null"`
	Operation string `gorm:"type:varchar(20);not null"`
}

func (o *OperationLog) TableName() string {
	return "operation_log"
}

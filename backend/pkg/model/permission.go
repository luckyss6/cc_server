package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	PermissionID   string `gorm:"type:varchar(64);uniqueIndex;not null"`
	PermissionName string `gorm:"type:varchar(20);uniqueIndex;not null"`
}
package model

import "gorm.io/gorm"

type MachineInfo struct {
	gorm.Model
	UserID   string `gorm:"type:varchar(64);not null"`
	Hostname string `gorm:"type:varchar(20);not null"`
	IP       string `gorm:"type:varchar(20);not null"`
	Host     string `gorm:"type:varchar(20);not null"`
	Username string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(20);not null"`
}

func (m *MachineInfo) TableName() string {
	return "machine_info"
}

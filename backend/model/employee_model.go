package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Email       string `gorm:"unique;not null" json:"email"`
	Telelephone int64  `gorm:"not null" json:"telephone"`
	Status      string `gorm:"type:varchar(20);check:status IN ('aktif', 'tidak aktif');default:'aktif';not null" json:"status"`
}

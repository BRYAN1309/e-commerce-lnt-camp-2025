package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name              string  `gorm:"type:varchar(150);not null" json:"name"`
	Description       string  `gorm:"type:text" json:"description"`
	PriceCents        int64   `gorm:"not null" json:"price_cents"` // store in cents to avoid float issues
	Stock             int     `gorm:"not null;default:0" json:"stock"`
	ImageURL          *string `gorm:"type:varchar(255)" json:"image_url"`
	ResponsibleUserID *uint   `json:"responsible_user_id"`
}

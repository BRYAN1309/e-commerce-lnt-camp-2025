package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"Description"`
	Price       float64 `json:"Price"`
	CreatedAt   time.Time
	Stock       int `json:"Stock"`
}

package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(150);not null"`
	Description string
	Price       float64 `gorm:"not null"`
	CategoryID  uint
	Category    Category
}

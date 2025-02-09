package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `gorm:"unique;not null"`
	Description string    `gorm:"type:text"`
	Products    []Product // One-to-Many relationship with Product
}

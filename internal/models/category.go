package models

type Category struct {
	BaseModel
	Name        string    `gorm:"unique;not null"`
	Description string    `gorm:"type:text"`
	Products    []Product // One-to-Many relationship with Product
}

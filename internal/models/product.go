package models

type Product struct {
	BaseModel
	Name        string `gorm:"type:varchar(150);not null"`
	Description string
	Price       float64 `gorm:"not null"`
	CategoryID  uint
	Category    Category
}

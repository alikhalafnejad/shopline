package models

type Product struct {
	BaseModel
	Name        string    `json:"name" gorm:"not null" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"omitempty,max=500"`
	Price       float64   `json:"price" gorm:"not null" validate:"required,gt=0"`
	CategoryID  uint      `json:"category_id" gorm:"not null"`
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
	Promoted    bool      `json:"promoted" gorm:"default:false"`
	Comments    []Comment `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

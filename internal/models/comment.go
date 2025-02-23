package models

import "shopline/pkg/constants"

type Comment struct {
	BaseModel
	UserID    uint             `json:"user_id" gorm:"not null"`
	ProductID uint             `json:"product_id" gorm:"not null"`
	Text      string           `json:"text" gorm:"not null" validate:"required,min=1,max=500"`
	Rating    int              `json:"rating" gorm:"not null" validate:"gte=1,lte=5"` // Rating between 1 and 5
	Status    constants.Status `json:"status" gorm:"type:varchar(20);default:pending" validate:"oneof=pending published rejected"`
}

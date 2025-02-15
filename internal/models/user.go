package models

type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique;not null" validate:"required,min=3,max=50"`
	Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required,min=6"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
}

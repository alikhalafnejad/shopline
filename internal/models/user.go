package models

type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique;not null" validate:"required,min=3,max=50"`
	Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required,min=6"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
	RoleID   uint   `json:"role_id" gorm:"not null" validate:"required,uint"`
	Role     Role   `json:"role" gorm:"foreignkey:RoleID"`
}

type Role struct {
	BaseModel
	Name        string `json:"name" gorm:"unique;not null" validate:"required,min=3,max=50"`
	Description string `json:"description" gorm:"not null" validate:"required,min=3,max=255"`
}

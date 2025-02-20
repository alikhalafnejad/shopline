package repositories

import "gorm.io/gorm"

type AdminRepository struct {
	BaseRepository
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		BaseRepository: *NewBaseRepository(db),
	}
}

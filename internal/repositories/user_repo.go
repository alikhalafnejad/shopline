package repositories

import (
	"gorm.io/gorm"
	"shopline/internal/models"
	"shopline/pkg/pagination"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository: *NewBaseRepository(db),
	}
}

// GetUserByEmail retrieves a user by their email address.
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers retrieves paginated users
func (r *UserRepository) GetUsers(page, limit int) (*pagination.PaginatedResponse, error) {
	query := r.DB
	var users []models.User
	return pagination.Paginate(query, page, limit, &users)
}

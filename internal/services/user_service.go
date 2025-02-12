package services

import (
	"errors"
	"shopline/internal/models"
	"shopline/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user *models.User) error {
	// Add validation or business logic here
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}
	return s.repo.Create(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user *models.User
	err := s.repo.FindByID(&user, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

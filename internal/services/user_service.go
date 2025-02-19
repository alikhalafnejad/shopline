package services

import (
	"errors"
	"shopline/internal/models"
	"shopline/internal/repositories"
	"shopline/pkg/auth"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(user *models.User) error {
	// Hashing the password before saving it in the database
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
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

// AuthenticateUser authenticates a user by email and password
func (s *UserService) AuthenticateUser(email, password string) (string, error) {
	// Fetch the user by email
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Verify the password
	if !auth.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	// Generate a JWT token
	token, err := auth.GenerateJWT(user)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}

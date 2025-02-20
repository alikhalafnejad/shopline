package services

import "shopline/internal/repositories"

type AdminService struct {
	repo *repositories.AdminRepository
}

func NewAdminService(repo *repositories.AdminRepository) *AdminService {
	return &AdminService{repo: repo}
}

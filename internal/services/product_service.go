package services

import (
	"shopline/internal/models"
	"shopline/internal/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(page, limit int) ([]models.Product, error) {
	return s.repo.GetProducts(page, limit)
}

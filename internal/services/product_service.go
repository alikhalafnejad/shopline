package services

import (
	"shopline/internal/models"
	"shopline/internal/repositories"
	"shopline/pkg/redisdb"
)

type ProductService struct {
	repo       *repositories.ProductRepository
	redisCache *redisdb.RedisClient
}

func NewProductService(repo *repositories.ProductRepository, redisCache *redisdb.RedisClient) *ProductService {
	return &ProductService{
		repo:       repo,
		redisCache: redisCache,
	}
}

// GetProducts retrieves paginated products with optional filters.
func (s *ProductService) GetProducts(page, limit, categoryID int, minPrice, maxPrice float64) ([]models.Product, error) {
	return s.repo.GetProducts(page, limit, categoryID, minPrice, maxPrice)
}

// GetProductByID retrieves a single product by id
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.repo.GetProductByID(id)
}

// CreateProduct create a new product.
func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

// DeleteProduct deletes a product
func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}

// GetPromotedProducts retrieves a list of promoted products
func (s *ProductService) GetPromotedProducts() ([]models.Product, error) {
	return s.repo.GetPromotedProducts()
}

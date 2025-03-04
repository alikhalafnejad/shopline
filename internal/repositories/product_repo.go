package repositories

import (
	"errors"
	"gorm.io/gorm"
	"shopline/internal/models"
	"shopline/pkg/pagination"
)

type ProductRepository struct {
	BaseRepository
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		BaseRepository: *NewBaseRepository(db),
	}
}

// GetProducts retrieves paginated products
func (r *ProductRepository) GetProducts(page, limit, categoryID int, minPrice, maxPrice float64) (*pagination.PaginatedResponse, error) {
	var products []models.Product
	query := r.DB

	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}
	if minPrice > 0 {
		query = query.Where("price = ?", minPrice)
	}
	if maxPrice > 0 {
		query = query.Where("price = ?", maxPrice)
	}

	return pagination.Paginate(query, page, limit, &products)
}

// GetProductByID retrieve a single product by ID
func (r *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.DB.Preload("Category").Preload("Seller").First(&product, id).Error
	if err != nil {
		return nil, errors.New("product not found")
	}
	return &product, nil
}

// CreateProduct creates a new product
func (r *ProductRepository) CreateProduct(product *models.Product) error {
	return r.DB.Create(product).Error
}

// UpdateProduct updates an existing product.
func (r *ProductRepository) UpdateProduct(id uint, updates map[string]interface{}) error {
	return r.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteProduct delete a product
func (r *ProductRepository) DeleteProduct(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}

// GetPromotedProducts retrieves a list of promoted products
func (r *ProductRepository) GetPromotedProducts(page, limit int) (*pagination.PaginatedResponse, error) {
	query := r.DB.Where("promoted = ?", true)
	var products []models.Product
	return pagination.Paginate(query, page, limit, &products)
}

package repositories

import (
	"gorm.io/gorm"
	"shopline/internal/models"
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
func (r *ProductRepository) GetProducts(page, limit int) ([]models.Product, error) {
	var products []models.Product
	offset := (page - 1) * limit
	err := r.DB.Offset(offset).Limit(limit).Find(&products).Error
	return products, err
}

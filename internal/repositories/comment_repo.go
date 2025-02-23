package repositories

import (
	"gorm.io/gorm"
	"shopline/internal/models"
	"shopline/pkg/constants"
	"shopline/pkg/pagination"
)

type CommentRepository struct {
	BaseRepository
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		BaseRepository: *NewBaseRepository(db),
	}
}

// GetPublishedCommentsByProduct retrieves paginated published comments for a specific product.
func (r *CommentRepository) GetPublishedCommentsByProduct(productID uint, page, limit int) (*pagination.PaginatedResponse, error) {
	var comments []models.Comment
	query := r.DB
	if productID != 0 {
		query = query.Where("product_id = ? AND status = ?", productID, constants.Published)
	}
	return pagination.Paginate(query, page, limit, &comments)
}

// CreateComment creates a new comment.
func (r *CommentRepository) CreateComment(comment *models.Comment) error {
	return r.DB.Create(comment).Error
}

// UpdateCommentStatus updates the status of a comment(admin-only action).
func (r *CommentRepository) UpdateCommentStatus(id uint, status constants.Status) error {
	return r.DB.Model(&models.Comment{}).Where("id = ?", id).Update("status", status).Error
}

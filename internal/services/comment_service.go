package services

import (
	"errors"
	"shopline/internal/models"
	"shopline/internal/repositories"
	"shopline/pkg/constants"
	"shopline/pkg/pagination"
)

type CommentService struct {
	repo *repositories.CommentRepository
}

func NewCommentService(repo *repositories.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

// GetPublishedCommentsByProduct retrieves comments for specific product.
func (s *CommentService) GetPublishedCommentsByProduct(productID uint, page, limit int) (*pagination.PaginatedResponse, error) {
	return s.repo.GetPublishedCommentsByProduct(productID, page, limit)
}

// CreateComment creates a new comment
func (s *CommentService) CreateComment(comment *models.Comment) error {
	if comment.Status != constants.Pending {
		comment.Status = constants.Pending // Default status is "pending"
	}
	return s.repo.CreateComment(comment)
}

// UpdateCommentStatus updates the status of a comment (admin-only action)
func (s *CommentService) UpdateCommentStatus(id uint, status constants.Status) error {
	if status != constants.Published && status != constants.REJECTED {
		return errors.New("invalid status: must be 'published' or 'rejected'")
	}
	return s.repo.UpdateCommentStatus(id, status)
}

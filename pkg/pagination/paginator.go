package pagination

import "gorm.io/gorm"

// PaginatedResponse defines the structure of a paginated response.
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	TotalCount int         `json:"total_count"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
}

// Paginate applies pagination to a GORM query and returns the paginated results.
func Paginate(db *gorm.DB, page, limit int, model interface{}) (*PaginatedResponse, error) {
	if page <= 0 {
		page = 1 // Default page to 1 if not provided
	}
	if limit <= 0 {
		limit = 10 // Default to 10 items per page if not provided
	}

	offset := (page - 1) * limit

	// Count the total number of records
	var totalCount int64
	if err := db.Model(model).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	// Apply pagination to the query
	query := db.Offset(offset).Limit(limit)
	if err := query.Find(model).Error; err != nil {
		return nil, err
	}

	return &PaginatedResponse{
		Data:       model,
		TotalCount: int(totalCount),
		Page:       page,
		Limit:      limit,
	}, nil
}

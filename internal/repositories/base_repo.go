package repositories

import "gorm.io/gorm"

type BaseRepository struct {
	DB *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{DB: db}
}

// Create inserts a new record into the database.
func (r *BaseRepository) Create(model interface{}) error {
	return r.DB.Create(model).Error
}

// FindByID retrieves a record by its ID.
func (r *BaseRepository) FindByID(model interface{}, id uint) error {
	return r.DB.First(model, id).Error
}

// Update updates a record in the database.
func (r *BaseRepository) Update(model interface{}, updates map[string]interface{}) error {
	return r.DB.Model(model).Updates(updates).Error
}

// Delete soft-deletes a record (if using GORM's soft delete feature).
func (r *BaseRepository) Delete(model interface{}, id uint) error {
	return r.DB.Delete(model, id).Error
}

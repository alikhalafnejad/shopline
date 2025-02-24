package repositories

import (
	"context"
	"gorm.io/gorm"
	"reflect"
	"shopline/pkg/cache"
	"time"
)

type BaseRepository struct {
	DB    *gorm.DB
	Cache *cache.CachedExecutor
}

func NewBaseRepository(db *gorm.DB, cacheClient cache.Client) *BaseRepository {
	return &BaseRepository{
		DB:    db,
		Cache: cache.NewCachedExecutor(cacheClient),
	}
}

type CachedFindOption func() []cache.CacheableFunc

func WithCache(key string, ttl time.Duration) CachedFindOption {
	return func() []cache.CacheableFunc {
		return []cache.CacheableFunc{
			cache.WithKey(key),
			cache.WithTTL(ttl),
		}
	}
}

func (r *BaseRepository) CachedFindOne(ctx context.Context, out interface{}, query interface{}, opts ...CachedFindOption) error {
	var cacheOpts []cache.CacheableFunc
	for _, opt := range opts {
		cacheOpts = append(cacheOpts, opt()...)
	}

	fn := func() (interface{}, error) {
		result := r.DB.Where(query).First(out)
		return result, result.Error
	}

	val, err := r.Cache.Execute(ctx, fn, cacheOpts...)
	if err != nil {
		return err
	}

	reflect.ValueOf(out).Elem().Set(reflect.ValueOf(val).Elem())
	return nil
}

func (r *BaseRepository) InvalidateCache(ctx context.Context, keys ...string) error {
	for _, key := range keys {
		if err := r.Cache.Delete(ctx, key); err != nil {
			return err
		}
	}
}

func (r *BaseRepository) CachedFind(ctx context.Context, query interface{}, opts ...CachedFindOption) (interface{}, error) {
	var cacheOpts []cache.CacheableFunc
	for _, opt := range opts {
		cacheOpts = append(cacheOpts, opt()...)
	}
	fn := func() (interface{}, error) {
		result := r.DB.Find(query)
		return result, result.Error
	}
	val, err := r.Cache.Execute(ctx, fn, cacheOpts...)
	if err != nil {
		return nil, err
	}
	reflect.ValueOf(query).Elem().Set(reflect.ValueOf(val).Elem())
	return val, nil
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

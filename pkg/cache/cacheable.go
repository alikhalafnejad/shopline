package cache

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"reflect"
	"shopline/pkg/logger"
	"time"
)

type CacheableOptions struct {
	KeyTemplate string
	TTL         time.Duration
	Disabled    bool
}

type CacheableFunc func(ctx context.Context, opts *CacheableOptions)

func WithKey(key string) CacheableFunc {
	return func(_ context.Context, opts *CacheableOptions) {
		opts.KeyTemplate = key
	}
}

func WithTTL(ttl time.Duration) CacheableFunc {
	return func(_ context.Context, opts *CacheableOptions) {
		opts.TTL = ttl
	}
}

type CachedExecutor struct {
	client Client
}

func NewCachedExecutor(client Client) *CachedExecutor {
	return &CachedExecutor{client: client}
}

func (c *CachedExecutor) Execute(ctx context.Context, fn interface{}, opts ...CacheableFunc) (interface{}, error) {
	config := &CacheableOptions{TTL: 5 * time.Minute}
	for _, opt := range opts {
		opt(ctx, config)
	}

	if config.Disabled || config.KeyTemplate == "" {
		return c.callFunction(fn)
	}

	cacheKey := c.generateCacheKey(fn, config.KeyTemplate)
	var result interface{}

	// Try to get from cache
	err := c.client.Get(cacheKey, &result)
	if err == nil {
		return result, nil
	}

	// Cache miss - execute function
	val, err := c.callFunction(fn)
	if err != nil {
		return nil, err
	}

	// Set to cache
	if err := c.client.Set(cacheKey, val, config.TTL); err != nil {
		logger.Logger.Error("Failed to set cache", zap.String("key", cacheKey), zap.Error(err))
	}

	return val, nil
}

func (c *CachedExecutor) generateCacheKey(fn interface{}, template string) string {
	fv := reflect.ValueOf(fn)
	args := make([]interface{}, fv.Type().NumIn())

	for i := 0; i < fv.Type().NumIn(); i++ {
		args[i] = fv.Index(i).Interface()
	}

	return fmt.Sprintf(template, args...)
}

func (c *CachedExecutor) callFunction(fn interface{}) (interface{}, error) {
	fv := reflect.ValueOf(fn)
	if fv.Kind() != reflect.Func {
		return nil, fmt.Errorf("non-function passed to Execute")
	}

	args := make([]reflect.Value, fv.Type().NumIn())
	for i := range args {
		if i == 0 {
			// Assume first argument is context.Context
			args[i] = reflect.ValueOf(context.Background())
		} else {
			// Initialize zero values for other arguments
			args[i] = reflect.Zero(fv.Type().In(i))
		}
	}

	results := fv.Call(args)
	if len(results) == 0 {
		return nil, nil
	}

	// Handle error as last return value
	var err error
	if len(results) > 1 {
		if e, ok := results[len(results)-1].Interface().(error); ok {
			err = e
		}
	}

	return results[0].Interface(), err
}

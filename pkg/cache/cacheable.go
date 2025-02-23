package cache

import (
	"context"
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

	cacheKey := c.
}

func (c *generate)

func (c *CachedExecutor) callFunction(fn interface{}) (interface{}, error) {

}

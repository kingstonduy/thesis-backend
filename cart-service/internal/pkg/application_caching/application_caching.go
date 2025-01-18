package application_caching

import (
	"context"
	"encoding/json"
	"fmt"

	cmap "github.com/orcaman/concurrent-map/v2"
)

type IClient interface {
	Get(ctx context.Context, key string, dest interface{}) error
	Set(ctx context.Context, key string, val interface{}) error
}

var (
	ErrKeyNotFound = fmt.Errorf("key is not found ")
)

type cachingRepository struct {
	m cmap.ConcurrentMap[string, []byte]
}

// NewCachingRepository creates a new caching repository.
func NewClient() IClient {
	return &cachingRepository{
		m: cmap.New[[]byte](),
	}
}

// Get implements IApplicationCachingRepo.
func (c *cachingRepository) Get(ctx context.Context, key string, dest interface{}) error {
	value, found := c.m.Get(key)

	if !found {
		return ErrKeyNotFound // Replace with your own error type for key not found
	}

	err := json.Unmarshal(value, &dest)

	return err
}

// Set implements IApplicationCachingRepo.
func (c *cachingRepository) Set(ctx context.Context, key string, value interface{}) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}

	c.m.Set(key, val)
	return nil
}

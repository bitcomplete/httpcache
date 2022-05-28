// Package redis provides a redis interface for http caching.
package redis

import (
	"context"
	"time"

	"github.com/bitcomplete/httpcache"
	"github.com/go-redis/redis/v8"
)

// cache is an implementation of httpcache.Cache that caches responses in a
// redis server.
type cache struct {
	client     *redis.Client
	expiration time.Duration
}

// cacheKey modifies an httpcache key for use in redis. Specifically, it
// prefixes keys to avoid collision with other data stored in redis.
func cacheKey(key string) string {
	return "rediscache:" + key
}

// Get returns the response corresponding to key if present.
func (c cache) Get(ctx context.Context, key string) (resp []byte, ok bool) {
	item, err := c.client.Get(ctx, cacheKey(key)).Bytes()
	if err != nil {
		return nil, false
	}
	return item, true
}

// Set saves a response to the cache as key.
func (c cache) Set(ctx context.Context, key string, resp []byte) {
	c.client.Set(ctx, cacheKey(key), resp, c.expiration)
}

// Delete removes the response with key from the cache.
func (c cache) Delete(ctx context.Context, key string) {
	c.client.Del(ctx, cacheKey(key))
}

// Opt is a configuration option for creating a new Redis cache
type Opt func(*cache)

// WithExpiration configures a Redis cache by setting its expiration
func WithExpiration(expiration time.Duration) Opt {
	return func(c *cache) {
		c.expiration = expiration
	}
}

// NewWithClient returns a new Cache with the given redis connection.
func NewWithClient(client *redis.Client, opts ...Opt) httpcache.Cache {
	c := cache{client: client}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

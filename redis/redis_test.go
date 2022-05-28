package redis

import (
	"context"
	"testing"

	"github.com/bitcomplete/httpcache/test"
	"github.com/go-redis/redis/v8"
)

func TestRedisCache(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		// TODO: rather than skip the test, fall back to a faked redis server
		t.Skipf("skipping test; no server running at localhost:6379")
	}
	client.FlushAll(context.Background())

	test.Cache(t, NewWithClient(client))
}

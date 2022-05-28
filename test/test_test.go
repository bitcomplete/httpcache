package test_test

import (
	"testing"

	"github.com/bitcomplete/httpcache"
	"github.com/bitcomplete/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}

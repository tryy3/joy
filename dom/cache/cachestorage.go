package cache

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/fetch"
)

type CacheStorage struct {
}

func (*CacheStorage) Delete(cacheName string) (b bool) {
	macro.Rewrite("await $_.delete($1)", cacheName)
	return b
}

func (*CacheStorage) Has(cacheName string) (b bool) {
	macro.Rewrite("await $_.has($1)", cacheName)
	return b
}

func (*CacheStorage) Keys() (s []string) {
	macro.Rewrite("await $_.keys()")
	return s
}

func (*CacheStorage) Match(request *fetch.Request, options *CacheQueryOptions) (i interface{}) {
	macro.Rewrite("await $_.match($1, $2)", request, options)
	return i
}

func (*CacheStorage) Open(cacheName string) (c *Cache) {
	macro.Rewrite("await $_.open($1)", cacheName)
	return c
}

package cache

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/fetch"
)

type Cache struct {
}

func (*Cache) Add(request *fetch.Request) {
	macro.Rewrite("await $_.add($1)", request)
}

func (*Cache) AddAll(requests []*fetch.Request) {
	macro.Rewrite("await $_.addAll($1)", requests)
}

func (*Cache) Delete(request *fetch.Request, options *CacheQueryOptions) (b bool) {
	macro.Rewrite("await $_.delete($1, $2)", request, options)
	return b
}

func (*Cache) Keys(request *fetch.Request, options *CacheQueryOptions) (r []*fetch.Request) {
	macro.Rewrite("await $_.keys($1, $2)", request, options)
	return r
}

func (*Cache) Match(request *fetch.Request, options *CacheQueryOptions) (r *fetch.Response) {
	macro.Rewrite("await $_.match($1, $2)", request, options)
	return r
}

func (*Cache) MatchAll(request *fetch.Request, options *CacheQueryOptions) (r []*fetch.Response) {
	macro.Rewrite("await $_.matchAll($1, $2)", request, options)
	return r
}

func (*Cache) Put(request *fetch.Request, response *fetch.Response) {
	macro.Rewrite("await $_.put($1, $2)", request, response)
}

package cache


import (
	"github.com/matthewmueller/joy/dom/fetch"
	"github.com/matthewmueller/joy/macro"
)

// Cache struct
// js:"Cache,omit"
type Cache struct {
}

// Add fn
// js:"add"
func (*Cache) Add(request *fetch.Request) {
	macro.Rewrite("await $_.add($1)", request)
}

// AddAll fn
// js:"addAll"
func (*Cache) AddAll(requests []*fetch.Request) {
	macro.Rewrite("await $_.addAll($1)", requests)
}

// Delete fn
// js:"delete"
func (*Cache) Delete(request *fetch.Request, options *CacheQueryOptions) (b bool) {
	macro.Rewrite("await $_.delete($1, $2)", request, options)
	return b
}

// Keys fn
// js:"keys"
func (*Cache) Keys(request *fetch.Request, options *CacheQueryOptions) (r []*fetch.Request) {
	macro.Rewrite("await $_.keys($1, $2)", request, options)
	return r
}

// Match fn
// js:"match"
func (*Cache) Match(request *fetch.Request, options *CacheQueryOptions) (r *fetch.Response) {
	macro.Rewrite("await $_.match($1, $2)", request, options)
	return r
}

// MatchAll fn
// js:"matchAll"
func (*Cache) MatchAll(request *fetch.Request, options *CacheQueryOptions) (r []*fetch.Response) {
	macro.Rewrite("await $_.matchAll($1, $2)", request, options)
	return r
}

// Put fn
// js:"put"
func (*Cache) Put(request *fetch.Request, response *fetch.Response) {
	macro.Rewrite("await $_.put($1, $2)", request, response)
}

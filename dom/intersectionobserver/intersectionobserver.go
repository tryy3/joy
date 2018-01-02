package intersectionobserver

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/element"
)

func New(callback func(entries []*utils.IntersectionObserverEntry, observer *IntersectionObserver), options *IntersectionObserverInit) *IntersectionObserver {
	macro.Rewrite("new IntersectionObserver($1, $2)", callback, options)
	return &IntersectionObserver{}
}

type IntersectionObserver struct {
}

func (*IntersectionObserver) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*IntersectionObserver) Observe(target element.Element) {
	macro.Rewrite("$_.observe($1)", target)
}

func (*IntersectionObserver) TakeRecords() (i []*utils.IntersectionObserverEntry) {
	macro.Rewrite("$_.takeRecords()")
	return i
}

func (*IntersectionObserver) Unobserve(target element.Element) {
	macro.Rewrite("$_.unobserve($1)", target)
}

func (*IntersectionObserver) Root() (root element.Element) {
	macro.Rewrite("$_.root")
	return root
}

func (*IntersectionObserver) RootMargin() (rootMargin string) {
	macro.Rewrite("$_.rootMargin")
	return rootMargin
}

func (*IntersectionObserver) Thresholds() (thresholds []float32) {
	macro.Rewrite("$_.thresholds")
	return thresholds
}

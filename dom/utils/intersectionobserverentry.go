package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/intersectionobserver"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/element"
)

func New(intersectionobserverentryinit *intersectionobserver.IntersectionObserverEntryInit) *IntersectionObserverEntry {
	macro.Rewrite("new IntersectionObserverEntry($1)", intersectionobserverentryinit)
	return &IntersectionObserverEntry{}
}

type IntersectionObserverEntry struct {
}

func (*IntersectionObserverEntry) BoundingClientRect() (boundingClientRect *dom.ClientRect) {
	macro.Rewrite("$_.boundingClientRect")
	return boundingClientRect
}

func (*IntersectionObserverEntry) IntersectionRatio() (intersectionRatio float32) {
	macro.Rewrite("$_.intersectionRatio")
	return intersectionRatio
}

func (*IntersectionObserverEntry) IntersectionRect() (intersectionRect *dom.ClientRect) {
	macro.Rewrite("$_.intersectionRect")
	return intersectionRect
}

func (*IntersectionObserverEntry) RootBounds() (rootBounds *dom.ClientRect) {
	macro.Rewrite("$_.rootBounds")
	return rootBounds
}

func (*IntersectionObserverEntry) Target() (target element.Element) {
	macro.Rewrite("$_.target")
	return target
}

func (*IntersectionObserverEntry) Time() (time int) {
	macro.Rewrite("$_.time")
	return time
}

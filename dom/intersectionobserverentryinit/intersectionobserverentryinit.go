package intersectionobserverentryinit

import (
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/window"
)

type IntersectionObserverEntryInit struct {
	boundingClientRect *dom.DOMRectInit
	intersectionRect   *dom.DOMRectInit
	rootBounds         *dom.DOMRectInit
	target             window.Element
	time               int
}

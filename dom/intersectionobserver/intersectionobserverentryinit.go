package intersectionobserver

import (
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/element"
)

type IntersectionObserverEntryInit struct {
	boundingClientRect	*dom.DOMRectInit
	intersectionRect	*dom.DOMRectInit
	rootBounds		*dom.DOMRectInit
	target			element.Element
	time			int
}

package intersectionobserver

import "github.com/matthewmueller/joy/dom/element"

type IntersectionObserverInit struct {
	root		*element.Element
	rootMargin	*string
	threshold	*interface{}
}

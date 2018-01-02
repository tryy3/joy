package html

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
)

var _ HTMLCollection = (*HTMLAreasCollection)(nil)

type HTMLAreasCollection struct {
}

func (*HTMLAreasCollection) Item(index uint) (w element.Element) {
	macro.Rewrite("$_.item($1)", index)
	return w
}

func (*HTMLAreasCollection) NamedItem(name string) (w element.Element) {
	macro.Rewrite("$_.namedItem($1)", name)
	return w
}

func (*HTMLAreasCollection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

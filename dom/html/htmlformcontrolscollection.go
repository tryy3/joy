package html

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
)

var _ HTMLCollection = (*HTMLFormControlsCollection)(nil)

type HTMLFormControlsCollection struct {
}

func (*HTMLFormControlsCollection) NamedItem(name string) (w element.Element) {
	macro.Rewrite("$_.namedItem($1)", name)
	return w
}

func (*HTMLFormControlsCollection) Item(index uint) (w element.Element) {
	macro.Rewrite("$_.item($1)", index)
	return w
}

func (*HTMLFormControlsCollection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

package html

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
)

var _ HTMLCollection = (*HTMLOptionsCollection)(nil)

type HTMLOptionsCollection struct {
}

func (*HTMLOptionsCollection) Add(element interface{}, before *interface{}) {
	macro.Rewrite("$_.add($1, $2)", element, before)
}

func (*HTMLOptionsCollection) Remove(index int) {
	macro.Rewrite("$_.remove($1)", index)
}

func (*HTMLOptionsCollection) Item(index uint) (w element.Element) {
	macro.Rewrite("$_.item($1)", index)
	return w
}

func (*HTMLOptionsCollection) NamedItem(name string) (w element.Element) {
	macro.Rewrite("$_.namedItem($1)", name)
	return w
}

func (*HTMLOptionsCollection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*HTMLOptionsCollection) SetLength(length uint) {
	macro.Rewrite("$_.length = $1", length)
}

func (*HTMLOptionsCollection) SelectedIndex() (selectedIndex int) {
	macro.Rewrite("$_.selectedIndex")
	return selectedIndex
}

func (*HTMLOptionsCollection) SetSelectedIndex(selectedIndex int) {
	macro.Rewrite("$_.selectedIndex = $1", selectedIndex)
}

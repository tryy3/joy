package dom

import "github.com/matthewmueller/joy/macro"

type SVGStringList struct {
}

func (*SVGStringList) AppendItem(newItem string) (s string) {
	macro.Rewrite("$_.appendItem($1)", newItem)
	return s
}

func (*SVGStringList) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*SVGStringList) GetItem(index uint) (s string) {
	macro.Rewrite("$_.getItem($1)", index)
	return s
}

func (*SVGStringList) Initialize(newItem string) (s string) {
	macro.Rewrite("$_.initialize($1)", newItem)
	return s
}

func (*SVGStringList) InsertItemBefore(newItem string, index uint) (s string) {
	macro.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGStringList) RemoveItem(index uint) (s string) {
	macro.Rewrite("$_.removeItem($1)", index)
	return s
}

func (*SVGStringList) ReplaceItem(newItem string, index uint) (s string) {
	macro.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGStringList) NumberOfItems() (numberOfItems uint) {
	macro.Rewrite("$_.numberOfItems")
	return numberOfItems
}

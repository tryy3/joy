package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGPathSegList struct {
}

func (*SVGPathSegList) AppendItem(newItem SVGPathSeg) (s SVGPathSeg) {
	macro.Rewrite("$_.appendItem($1)", newItem)
	return s
}

func (*SVGPathSegList) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*SVGPathSegList) GetItem(index uint) (s SVGPathSeg) {
	macro.Rewrite("$_.getItem($1)", index)
	return s
}

func (*SVGPathSegList) Initialize(newItem SVGPathSeg) (s SVGPathSeg) {
	macro.Rewrite("$_.initialize($1)", newItem)
	return s
}

func (*SVGPathSegList) InsertItemBefore(newItem SVGPathSeg, index uint) (s SVGPathSeg) {
	macro.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGPathSegList) RemoveItem(index uint) (s SVGPathSeg) {
	macro.Rewrite("$_.removeItem($1)", index)
	return s
}

func (*SVGPathSegList) ReplaceItem(newItem SVGPathSeg, index uint) (s SVGPathSeg) {
	macro.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGPathSegList) NumberOfItems() (numberOfItems uint) {
	macro.Rewrite("$_.numberOfItems")
	return numberOfItems
}

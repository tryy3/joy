package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGLengthList struct {
}

func (*SVGLengthList) AppendItem(newItem *SVGLength) (s *SVGLength) {
	macro.Rewrite("$_.appendItem($1)", newItem)
	return s
}

func (*SVGLengthList) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*SVGLengthList) GetItem(index uint) (s *SVGLength) {
	macro.Rewrite("$_.getItem($1)", index)
	return s
}

func (*SVGLengthList) Initialize(newItem *SVGLength) (s *SVGLength) {
	macro.Rewrite("$_.initialize($1)", newItem)
	return s
}

func (*SVGLengthList) InsertItemBefore(newItem *SVGLength, index uint) (s *SVGLength) {
	macro.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGLengthList) RemoveItem(index uint) (s *SVGLength) {
	macro.Rewrite("$_.removeItem($1)", index)
	return s
}

func (*SVGLengthList) ReplaceItem(newItem *SVGLength, index uint) (s *SVGLength) {
	macro.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGLengthList) NumberOfItems() (numberOfItems uint) {
	macro.Rewrite("$_.numberOfItems")
	return numberOfItems
}

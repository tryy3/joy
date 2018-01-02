package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGNumberList struct {
}

func (*SVGNumberList) AppendItem(newItem *SVGNumber) (s *SVGNumber) {
	macro.Rewrite("$_.appendItem($1)", newItem)
	return s
}

func (*SVGNumberList) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*SVGNumberList) GetItem(index uint) (s *SVGNumber) {
	macro.Rewrite("$_.getItem($1)", index)
	return s
}

func (*SVGNumberList) Initialize(newItem *SVGNumber) (s *SVGNumber) {
	macro.Rewrite("$_.initialize($1)", newItem)
	return s
}

func (*SVGNumberList) InsertItemBefore(newItem *SVGNumber, index uint) (s *SVGNumber) {
	macro.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGNumberList) RemoveItem(index uint) (s *SVGNumber) {
	macro.Rewrite("$_.removeItem($1)", index)
	return s
}

func (*SVGNumberList) ReplaceItem(newItem *SVGNumber, index uint) (s *SVGNumber) {
	macro.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGNumberList) NumberOfItems() (numberOfItems uint) {
	macro.Rewrite("$_.numberOfItems")
	return numberOfItems
}

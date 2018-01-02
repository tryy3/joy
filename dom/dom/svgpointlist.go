package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGPointList struct {
}

func (*SVGPointList) AppendItem(newItem *SVGPoint) (s *SVGPoint) {
	macro.Rewrite("$_.appendItem($1)", newItem)
	return s
}

func (*SVGPointList) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*SVGPointList) GetItem(index uint) (s *SVGPoint) {
	macro.Rewrite("$_.getItem($1)", index)
	return s
}

func (*SVGPointList) Initialize(newItem *SVGPoint) (s *SVGPoint) {
	macro.Rewrite("$_.initialize($1)", newItem)
	return s
}

func (*SVGPointList) InsertItemBefore(newItem *SVGPoint, index uint) (s *SVGPoint) {
	macro.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGPointList) RemoveItem(index uint) (s *SVGPoint) {
	macro.Rewrite("$_.removeItem($1)", index)
	return s
}

func (*SVGPointList) ReplaceItem(newItem *SVGPoint, index uint) (s *SVGPoint) {
	macro.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGPointList) NumberOfItems() (numberOfItems uint) {
	macro.Rewrite("$_.numberOfItems")
	return numberOfItems
}

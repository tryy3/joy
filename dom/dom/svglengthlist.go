package dom

import "github.com/matthewmueller/joy/macro"

// SVGLengthList struct
// js:"SVGLengthList,omit"
type SVGLengthList struct {
}

// AppendItem fn
// js:"appendItem"
func (*SVGLengthList) AppendItem(newItem *SVGLength) (s *SVGLength) {
	macro.Rewrite("$_.appendItem($1)", newItem)
	return s
}

// Clear fn
// js:"clear"
func (*SVGLengthList) Clear() {
	macro.Rewrite("$_.clear()")
}

// GetItem fn
// js:"getItem"
func (*SVGLengthList) GetItem(index uint) (s *SVGLength) {
	macro.Rewrite("$_.getItem($1)", index)
	return s
}

// Initialize fn
// js:"initialize"
func (*SVGLengthList) Initialize(newItem *SVGLength) (s *SVGLength) {
	macro.Rewrite("$_.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
// js:"insertItemBefore"
func (*SVGLengthList) InsertItemBefore(newItem *SVGLength, index uint) (s *SVGLength) {
	macro.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
// js:"removeItem"
func (*SVGLengthList) RemoveItem(index uint) (s *SVGLength) {
	macro.Rewrite("$_.removeItem($1)", index)
	return s
}

// ReplaceItem fn
// js:"replaceItem"
func (*SVGLengthList) ReplaceItem(newItem *SVGLength, index uint) (s *SVGLength) {
	macro.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
// js:"numberOfItems"
func (*SVGLengthList) NumberOfItems() (numberOfItems uint) {
	macro.Rewrite("$_.numberOfItems")
	return numberOfItems
}

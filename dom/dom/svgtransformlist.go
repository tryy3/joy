package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGTransformList struct {
}

func (*SVGTransformList) AppendItem(newItem *SVGTransform) (s *SVGTransform) {
	macro.Rewrite("$_.appendItem($1)", newItem)
	return s
}

func (*SVGTransformList) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*SVGTransformList) Consolidate() (s *SVGTransform) {
	macro.Rewrite("$_.consolidate()")
	return s
}

func (*SVGTransformList) CreateSVGTransformFromMatrix(matrix *SVGMatrix) (s *SVGTransform) {
	macro.Rewrite("$_.createSVGTransformFromMatrix($1)", matrix)
	return s
}

func (*SVGTransformList) GetItem(index uint) (s *SVGTransform) {
	macro.Rewrite("$_.getItem($1)", index)
	return s
}

func (*SVGTransformList) Initialize(newItem *SVGTransform) (s *SVGTransform) {
	macro.Rewrite("$_.initialize($1)", newItem)
	return s
}

func (*SVGTransformList) InsertItemBefore(newItem *SVGTransform, index uint) (s *SVGTransform) {
	macro.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGTransformList) RemoveItem(index uint) (s *SVGTransform) {
	macro.Rewrite("$_.removeItem($1)", index)
	return s
}

func (*SVGTransformList) ReplaceItem(newItem *SVGTransform, index uint) (s *SVGTransform) {
	macro.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGTransformList) NumberOfItems() (numberOfItems uint) {
	macro.Rewrite("$_.numberOfItems")
	return numberOfItems
}

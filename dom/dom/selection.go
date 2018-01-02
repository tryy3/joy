package dom

import "github.com/matthewmueller/joy/macro"

type Selection struct {
}

func (*Selection) AddRange(rng *Range) {
	macro.Rewrite("$_.addRange($1)", rng)
}

func (*Selection) Collapse(parentNode Node, offset int) {
	macro.Rewrite("$_.collapse($1, $2)", parentNode, offset)
}

func (*Selection) CollapseToEnd() {
	macro.Rewrite("$_.collapseToEnd()")
}

func (*Selection) CollapseToStart() {
	macro.Rewrite("$_.collapseToStart()")
}

func (*Selection) ContainsNode(node Node, partlyContained bool) (b bool) {
	macro.Rewrite("$_.containsNode($1, $2)", node, partlyContained)
	return b
}

func (*Selection) DeleteFromDocument() {
	macro.Rewrite("$_.deleteFromDocument()")
}

func (*Selection) Empty() {
	macro.Rewrite("$_.empty()")
}

func (*Selection) Extend(newNode Node, offset int) {
	macro.Rewrite("$_.extend($1, $2)", newNode, offset)
}

func (*Selection) GetRangeAt(index int) (r *Range) {
	macro.Rewrite("$_.getRangeAt($1)", index)
	return r
}

func (*Selection) RemoveAllRanges() {
	macro.Rewrite("$_.removeAllRanges()")
}

func (*Selection) RemoveRange(rng *Range) {
	macro.Rewrite("$_.removeRange($1)", rng)
}

func (*Selection) SelectAllChildren(parentNode Node) {
	macro.Rewrite("$_.selectAllChildren($1)", parentNode)
}

func (*Selection) SetBaseAndExtent(baseNode Node, baseOffset int, extentNode Node, extentOffset int) {
	macro.Rewrite("$_.setBaseAndExtent($1, $2, $3, $4)", baseNode, baseOffset, extentNode, extentOffset)
}

func (*Selection) SetPosition(parentNode Node, offset int) {
	macro.Rewrite("$_.setPosition($1, $2)", parentNode, offset)
}

func (*Selection) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

func (*Selection) AnchorNode() (anchorNode Node) {
	macro.Rewrite("$_.anchorNode")
	return anchorNode
}

func (*Selection) AnchorOffset() (anchorOffset int) {
	macro.Rewrite("$_.anchorOffset")
	return anchorOffset
}

func (*Selection) BaseNode() (baseNode Node) {
	macro.Rewrite("$_.baseNode")
	return baseNode
}

func (*Selection) BaseOffset() (baseOffset int) {
	macro.Rewrite("$_.baseOffset")
	return baseOffset
}

func (*Selection) ExtentNode() (extentNode Node) {
	macro.Rewrite("$_.extentNode")
	return extentNode
}

func (*Selection) ExtentOffset() (extentOffset int) {
	macro.Rewrite("$_.extentOffset")
	return extentOffset
}

func (*Selection) FocusNode() (focusNode Node) {
	macro.Rewrite("$_.focusNode")
	return focusNode
}

func (*Selection) FocusOffset() (focusOffset int) {
	macro.Rewrite("$_.focusOffset")
	return focusOffset
}

func (*Selection) IsCollapsed() (isCollapsed bool) {
	macro.Rewrite("$_.isCollapsed")
	return isCollapsed
}

func (*Selection) RangeCount() (rangeCount int) {
	macro.Rewrite("$_.rangeCount")
	return rangeCount
}

func (*Selection) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

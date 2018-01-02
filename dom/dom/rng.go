package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/document"
)

type Range struct {
}

func (*Range) CloneContents() (d *document.DocumentFragment) {
	macro.Rewrite("$_.cloneContents()")
	return d
}

func (*Range) CloneRange() (r *Range) {
	macro.Rewrite("$_.cloneRange()")
	return r
}

func (*Range) Collapse(toStart bool) {
	macro.Rewrite("$_.collapse($1)", toStart)
}

func (*Range) CompareBoundaryPoints(how uint8, sourceRange *Range) (i int8) {
	macro.Rewrite("$_.compareBoundaryPoints($1, $2)", how, sourceRange)
	return i
}

func (*Range) CreateContextualFragment(fragment string) (d *document.DocumentFragment) {
	macro.Rewrite("$_.createContextualFragment($1)", fragment)
	return d
}

func (*Range) DeleteContents() {
	macro.Rewrite("$_.deleteContents()")
}

func (*Range) Detach() {
	macro.Rewrite("$_.detach()")
}

func (*Range) Expand(Unit *ExpandGranularity) (b bool) {
	macro.Rewrite("$_.expand($1)", Unit)
	return b
}

func (*Range) ExtractContents() (d *document.DocumentFragment) {
	macro.Rewrite("$_.extractContents()")
	return d
}

func (*Range) GetBoundingClientRect() (c *ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

func (*Range) GetClientRects() (c *ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

func (*Range) InsertNode(newNode Node) {
	macro.Rewrite("$_.insertNode($1)", newNode)
}

func (*Range) SelectNode(refNode Node) {
	macro.Rewrite("$_.selectNode($1)", refNode)
}

func (*Range) SelectNodeContents(refNode Node) {
	macro.Rewrite("$_.selectNodeContents($1)", refNode)
}

func (*Range) SetEnd(refNode Node, offset int) {
	macro.Rewrite("$_.setEnd($1, $2)", refNode, offset)
}

func (*Range) SetEndAfter(refNode Node) {
	macro.Rewrite("$_.setEndAfter($1)", refNode)
}

func (*Range) SetEndBefore(refNode Node) {
	macro.Rewrite("$_.setEndBefore($1)", refNode)
}

func (*Range) SetStart(refNode Node, offset int) {
	macro.Rewrite("$_.setStart($1, $2)", refNode, offset)
}

func (*Range) SetStartAfter(refNode Node) {
	macro.Rewrite("$_.setStartAfter($1)", refNode)
}

func (*Range) SetStartBefore(refNode Node) {
	macro.Rewrite("$_.setStartBefore($1)", refNode)
}

func (*Range) SurroundContents(newParent Node) {
	macro.Rewrite("$_.surroundContents($1)", newParent)
}

func (*Range) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

func (*Range) Collapsed() (collapsed bool) {
	macro.Rewrite("$_.collapsed")
	return collapsed
}

func (*Range) CommonAncestorContainer() (commonAncestorContainer Node) {
	macro.Rewrite("$_.commonAncestorContainer")
	return commonAncestorContainer
}

func (*Range) EndContainer() (endContainer Node) {
	macro.Rewrite("$_.endContainer")
	return endContainer
}

func (*Range) EndOffset() (endOffset int) {
	macro.Rewrite("$_.endOffset")
	return endOffset
}

func (*Range) StartContainer() (startContainer Node) {
	macro.Rewrite("$_.startContainer")
	return startContainer
}

func (*Range) StartOffset() (startOffset int) {
	macro.Rewrite("$_.startOffset")
	return startOffset
}

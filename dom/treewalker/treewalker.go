package treewalker

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
)

type TreeWalker struct {
}

func (*TreeWalker) FirstChild() (n dom.Node) {
	macro.Rewrite("$_.firstChild()")
	return n
}

func (*TreeWalker) LastChild() (n dom.Node) {
	macro.Rewrite("$_.lastChild()")
	return n
}

func (*TreeWalker) NextNode() (n dom.Node) {
	macro.Rewrite("$_.nextNode()")
	return n
}

func (*TreeWalker) NextSibling() (n dom.Node) {
	macro.Rewrite("$_.nextSibling()")
	return n
}

func (*TreeWalker) ParentNode() (n dom.Node) {
	macro.Rewrite("$_.parentNode()")
	return n
}

func (*TreeWalker) PreviousNode() (n dom.Node) {
	macro.Rewrite("$_.previousNode()")
	return n
}

func (*TreeWalker) PreviousSibling() (n dom.Node) {
	macro.Rewrite("$_.previousSibling()")
	return n
}

func (*TreeWalker) CurrentNode() (currentNode dom.Node) {
	macro.Rewrite("$_.currentNode")
	return currentNode
}

func (*TreeWalker) SetCurrentNode(currentNode dom.Node) {
	macro.Rewrite("$_.currentNode = $1", currentNode)
}

func (*TreeWalker) ExpandEntityReferences() (expandEntityReferences bool) {
	macro.Rewrite("$_.expandEntityReferences")
	return expandEntityReferences
}

func (*TreeWalker) Filter() (filter *dom.NodeFilter) {
	macro.Rewrite("$_.filter")
	return filter
}

func (*TreeWalker) Root() (root dom.Node) {
	macro.Rewrite("$_.root")
	return root
}

func (*TreeWalker) WhatToShow() (whatToShow uint) {
	macro.Rewrite("$_.whatToShow")
	return whatToShow
}

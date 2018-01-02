package dom

import "github.com/matthewmueller/joy/macro"

type NodeIterator struct {
}

func (*NodeIterator) Detach() {
	macro.Rewrite("$_.detach()")
}

func (*NodeIterator) NextNode() (n Node) {
	macro.Rewrite("$_.nextNode()")
	return n
}

func (*NodeIterator) PreviousNode() (n Node) {
	macro.Rewrite("$_.previousNode()")
	return n
}

func (*NodeIterator) ExpandEntityReferences() (expandEntityReferences bool) {
	macro.Rewrite("$_.expandEntityReferences")
	return expandEntityReferences
}

func (*NodeIterator) Filter() (filter *NodeFilter) {
	macro.Rewrite("$_.filter")
	return filter
}

func (*NodeIterator) Root() (root Node) {
	macro.Rewrite("$_.root")
	return root
}

func (*NodeIterator) WhatToShow() (whatToShow uint) {
	macro.Rewrite("$_.whatToShow")
	return whatToShow
}

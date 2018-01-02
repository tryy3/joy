package mutation

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
)

type MutationRecord struct {
}

func (*MutationRecord) AddedNodes() (addedNodes *dom.NodeList) {
	macro.Rewrite("$_.addedNodes")
	return addedNodes
}

func (*MutationRecord) AttributeName() (attributeName string) {
	macro.Rewrite("$_.attributeName")
	return attributeName
}

func (*MutationRecord) AttributeNamespace() (attributeNamespace string) {
	macro.Rewrite("$_.attributeNamespace")
	return attributeNamespace
}

func (*MutationRecord) NextSibling() (nextSibling dom.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*MutationRecord) OldValue() (oldValue string) {
	macro.Rewrite("$_.oldValue")
	return oldValue
}

func (*MutationRecord) PreviousSibling() (previousSibling dom.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*MutationRecord) RemovedNodes() (removedNodes *dom.NodeList) {
	macro.Rewrite("$_.removedNodes")
	return removedNodes
}

func (*MutationRecord) Target() (target dom.Node) {
	macro.Rewrite("$_.target")
	return target
}

func (*MutationRecord) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

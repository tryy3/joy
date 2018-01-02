package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
)

type XPathResult struct {
}

func (*XPathResult) IterateNext() (n dom.Node) {
	macro.Rewrite("$_.iterateNext()")
	return n
}

func (*XPathResult) SnapshotItem(index uint) (n dom.Node) {
	macro.Rewrite("$_.snapshotItem($1)", index)
	return n
}

func (*XPathResult) BooleanValue() (booleanValue bool) {
	macro.Rewrite("$_.booleanValue")
	return booleanValue
}

func (*XPathResult) InvalidIteratorState() (invalidIteratorState bool) {
	macro.Rewrite("$_.invalidIteratorState")
	return invalidIteratorState
}

func (*XPathResult) NumberValue() (numberValue float32) {
	macro.Rewrite("$_.numberValue")
	return numberValue
}

func (*XPathResult) ResultType() (resultType uint8) {
	macro.Rewrite("$_.resultType")
	return resultType
}

func (*XPathResult) SingleNodeValue() (singleNodeValue dom.Node) {
	macro.Rewrite("$_.singleNodeValue")
	return singleNodeValue
}

func (*XPathResult) SnapshotLength() (snapshotLength uint) {
	macro.Rewrite("$_.snapshotLength")
	return snapshotLength
}

func (*XPathResult) StringValue() (stringValue string) {
	macro.Rewrite("$_.stringValue")
	return stringValue
}

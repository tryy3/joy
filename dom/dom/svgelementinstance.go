package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*SVGElementInstance)(nil)

type SVGElementInstance struct {
}

func (*SVGElementInstance) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SVGElementInstance) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*SVGElementInstance) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SVGElementInstance) ChildNodes() (childNodes *SVGElementInstanceList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*SVGElementInstance) CorrespondingElement() (correspondingElement SVGElement) {
	macro.Rewrite("$_.correspondingElement")
	return correspondingElement
}

func (*SVGElementInstance) CorrespondingUseElement() (correspondingUseElement *SVGUseElement) {
	macro.Rewrite("$_.correspondingUseElement")
	return correspondingUseElement
}

func (*SVGElementInstance) FirstChild() (firstChild *SVGElementInstance) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*SVGElementInstance) LastChild() (lastChild *SVGElementInstance) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*SVGElementInstance) NextSibling() (nextSibling *SVGElementInstance) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*SVGElementInstance) ParentNode() (parentNode *SVGElementInstance) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*SVGElementInstance) PreviousSibling() (previousSibling *SVGElementInstance) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

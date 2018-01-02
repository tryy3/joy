package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
)

func New() *MSGesture {
	macro.Rewrite("new MSGesture()")
	return &MSGesture{}
}

type MSGesture struct {
}

func (*MSGesture) AddPointer(pointerId int) {
	macro.Rewrite("$_.addPointer($1)", pointerId)
}

func (*MSGesture) Stop() {
	macro.Rewrite("$_.stop()")
}

func (*MSGesture) Target() (target element.Element) {
	macro.Rewrite("$_.target")
	return target
}

func (*MSGesture) SetTarget(target element.Element) {
	macro.Rewrite("$_.target = $1", target)
}

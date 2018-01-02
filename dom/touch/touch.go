package touch

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

type Touch struct {
}

func (*Touch) ClientX() (clientX int) {
	macro.Rewrite("$_.clientX")
	return clientX
}

func (*Touch) ClientY() (clientY int) {
	macro.Rewrite("$_.clientY")
	return clientY
}

func (*Touch) Identifier() (identifier int) {
	macro.Rewrite("$_.identifier")
	return identifier
}

func (*Touch) PageX() (pageX int) {
	macro.Rewrite("$_.pageX")
	return pageX
}

func (*Touch) PageY() (pageY int) {
	macro.Rewrite("$_.pageY")
	return pageY
}

func (*Touch) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

func (*Touch) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

func (*Touch) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

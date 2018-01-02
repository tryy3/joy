package gamepad

import (
	"github.com/matthewmueller/joy/macro"
)

type Gamepad struct {
}

func (*Gamepad) Axes() (axes []float32) {
	macro.Rewrite("$_.axes")
	return axes
}

func (*Gamepad) Buttons() (buttons []*GamepadButton) {
	macro.Rewrite("$_.buttons")
	return buttons
}

func (*Gamepad) Connected() (connected bool) {
	macro.Rewrite("$_.connected")
	return connected
}

func (*Gamepad) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*Gamepad) Index() (index int) {
	macro.Rewrite("$_.index")
	return index
}

func (*Gamepad) Mapping() (mapping string) {
	macro.Rewrite("$_.mapping")
	return mapping
}

func (*Gamepad) Timestamp() (timestamp int) {
	macro.Rewrite("$_.timestamp")
	return timestamp
}

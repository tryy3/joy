package gamepad

import "github.com/matthewmueller/joy/macro"

type GamepadButton struct {
}

func (*GamepadButton) Pressed() (pressed bool) {
	macro.Rewrite("$_.pressed")
	return pressed
}

func (*GamepadButton) Value() (value float32) {
	macro.Rewrite("$_.value")
	return value
}

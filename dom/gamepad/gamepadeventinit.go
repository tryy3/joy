package gamepad

import "github.com/matthewmueller/joy/dom/event"

type GamepadEventInit struct {
	*event.EventInit

	gamepad	*Gamepad
}

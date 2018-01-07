package gamepadeventinit

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/gamepad"
)

type GamepadEventInit struct {
	*event.EventInit

	gamepad *gamepad.Gamepad
}

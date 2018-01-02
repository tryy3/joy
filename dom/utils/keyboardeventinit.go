package utils

import "github.com/matthewmueller/joy/dom/event"

type KeyboardEventInit struct {
	*event.EventModifierInit

	key		*string
	location	*uint
	repeat		*bool
}

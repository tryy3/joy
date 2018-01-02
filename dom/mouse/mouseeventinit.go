package mouse

import "github.com/matthewmueller/joy/dom/event"

type MouseEventInit struct {
	*event.EventModifierInit

	button		*int8
	buttons		*uint8
	clientX		*int
	clientY		*int
	relatedTarget	*event.EventTarget
	screenX		*int
	screenY		*int
}

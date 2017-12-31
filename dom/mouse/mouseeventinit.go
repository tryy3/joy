package mouse


import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/ui"
)

type MouseEventInit struct {
	*ui.EventModifierInit

	button        *int8
	buttons       *uint8
	clientX       *int
	clientY       *int
	relatedTarget *event.EventTarget
	screenX       *int
	screenY       *int
}

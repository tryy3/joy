package utils

import (
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
)

type FocusEventInit struct {
	*ui.UIEventInit

	relatedTarget	*event.EventTarget
}

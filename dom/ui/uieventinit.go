package ui

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
)

type UIEventInit struct {
	*event.EventInit

	detail	*int
	view	*window.Window
}

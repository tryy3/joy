package ui

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
)

type UIEvent interface {
	event.
		Event

	InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int)

	Detail() (detail int)

	View() (view *window.Window)
}

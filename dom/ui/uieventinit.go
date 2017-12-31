package ui


import "github.com/matthewmueller/joy/dom/event"

type UIEventInit struct {
	*event.EventInit

	detail *int
	view   *Window
}

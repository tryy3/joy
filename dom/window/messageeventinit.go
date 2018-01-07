package window

import "github.com/matthewmueller/joy/dom/event"

type MessageEventInit struct {
	*event.EventInit

	data   *interface{}
	origin *string
	ports  *[]*MessagePort
	source *Window
}

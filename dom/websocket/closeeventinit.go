package websocket

import "github.com/matthewmueller/joy/dom/event"

type CloseEventInit struct {
	*event.EventInit

	code     *uint8
	reason   *string
	wasClean *bool
}

package mediastreameventinit

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
)

type MediaStreamEventInit struct {
	*event.EventInit

	stream *window.MediaStream
}

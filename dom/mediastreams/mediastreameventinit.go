package mediastreams

import "github.com/matthewmueller/joy/dom/event"

type MediaStreamEventInit struct {
	*event.EventInit

	stream	*MediaStream
}

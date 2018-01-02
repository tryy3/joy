package mediastreams

import "github.com/matthewmueller/joy/dom/event"

type MediaStreamErrorEventInit struct {
	*event.EventInit

	err	*MediaStreamError
}

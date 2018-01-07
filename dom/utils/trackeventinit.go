package utils

import "github.com/matthewmueller/joy/dom/event"

type TrackEventInit struct {
	*event.EventInit

	track *interface{}
}

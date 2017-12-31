package utils


import "github.com/matthewmueller/joy/dom/eventinit"

type TrackEventInit struct {
	*eventinit.EventInit

	track *interface{}
}

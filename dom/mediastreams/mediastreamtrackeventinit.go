package mediastreams

import "github.com/matthewmueller/joy/dom/event"

type MediaStreamTrackEventInit struct {
	*event.EventInit

	track	*MediaStreamTrack
}

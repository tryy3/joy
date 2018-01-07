package window

import "github.com/matthewmueller/joy/dom/mediastreams"

// NavigatorUserMedia interface
// js:"NavigatorUserMedia"
type NavigatorUserMedia interface {

	// GetUserMedia
	// js:"getUserMedia"
	// jsrewrite:"$_.getUserMedia($1, $2, $3)"
	GetUserMedia(constraints *mediastreams.MediaStreamConstraints, successCallback func(stream *MediaStream), errorCallback func(err *mediastreams.MediaStreamError))

	// MediaDevices prop
	// js:"mediaDevices"
	// jsrewrite:"$_.mediaDevices"
	MediaDevices() (mediaDevices *MediaDevices)
}

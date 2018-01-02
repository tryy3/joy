package window

import "github.com/matthewmueller/joy/dom/mediastreams"

type NavigatorUserMedia interface {
	GetUserMedia(constraints *mediastreams.MediaStreamConstraints, successCallback func(stream *mediastreams.MediaStream), errorCallback func(err *mediastreams.MediaStreamError))

	MediaDevices() (mediaDevices *mediastreams.MediaDevices)
}

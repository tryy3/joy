package webrtc

import "github.com/matthewmueller/joy/dom/event"

type RTCDTMFToneChangeEventInit struct {
	*event.EventInit

	tone	*string
}

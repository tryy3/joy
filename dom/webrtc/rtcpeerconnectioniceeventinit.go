package webrtc

import "github.com/matthewmueller/joy/dom/event"

type RTCPeerConnectionIceEventInit struct {
	*event.EventInit

	candidate	*RTCIceCandidate
}

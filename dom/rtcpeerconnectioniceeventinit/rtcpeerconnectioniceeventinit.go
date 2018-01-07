package rtcpeerconnectioniceeventinit

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/webrtc"
)

type RTCPeerConnectionIceEventInit struct {
	*event.EventInit

	candidate *webrtc.RTCIceCandidate
}

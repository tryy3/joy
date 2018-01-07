package webrtc

import "github.com/matthewmueller/joy/dom/ms"

type RTCIceGatherOptions struct {
	gatherPolicy *RTCIceGatherPolicy
	iceservers   *[]*RTCIceServer
	portRange    *ms.MSPortRange
}

package rtcicecandidatepair

import "github.com/matthewmueller/joy/dom/webrtc"

type RTCIceCandidatePair struct {
	local  *webrtc.RTCIceCandidateDictionary
	remote *webrtc.RTCIceCandidateDictionary
}

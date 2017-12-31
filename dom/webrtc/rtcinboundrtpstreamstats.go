package webrtc


import "github.com/matthewmueller/joy/dom/rtcrtpstreamstats"

type RTCInboundRTPStreamStats struct {
	*rtcrtpstreamstats.RTCRTPStreamStats

	bytesReceived   *uint64
	fractionLost    *float32
	jitter          *float32
	packetsLost     *uint
	packetsReceived *uint
}

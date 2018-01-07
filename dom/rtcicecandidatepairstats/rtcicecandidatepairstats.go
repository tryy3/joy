package rtcicecandidatepairstats

import "github.com/matthewmueller/joy/dom/webrtc"

type RTCIceCandidatePairStats struct {
	*webrtc.RTCStats

	availableIncomingBitrate *float32
	availableOutgoingBitrate *float32
	bytesReceived            *uint64
	bytesSent                *uint64
	localCandidateId         *string
	nominated                *bool
	priority                 *uint64
	readable                 *bool
	remoteCandidateId        *string
	roundTripTime            *float32
	state                    *webrtc.RTCStatsIceCandidatePairState
	transportId              *string
	writable                 *bool
}

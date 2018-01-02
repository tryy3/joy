package ms

import "github.com/matthewmueller/joy/dom/webrtc"

type MSNetwork struct {
	*webrtc.RTCStats

	delay		*MSDelay
	jitter		*MSJitter
	packetLoss	*MSPacketLoss
	utilization	*MSUtilization
}

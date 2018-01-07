package msnetwork

import (
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/webrtc"
)

type MSNetwork struct {
	*webrtc.RTCStats

	delay       *ms.MSDelay
	jitter      *ms.MSJitter
	packetLoss  *ms.MSPacketLoss
	utilization *ms.MSUtilization
}

package rtcicecandidateattributes

import "github.com/matthewmueller/joy/dom/webrtc"

type RTCIceCandidateAttributes struct {
	*webrtc.RTCStats

	addressSourceUrl *string
	candidateType    *webrtc.RTCStatsIceCandidateType
	ipAddress        *string
	portNumber       *int
	priority         *int
	transport        *string
}

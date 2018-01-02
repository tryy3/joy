package webrtc

type RTCIceCandidateAttributes struct {
	*RTCStats

	addressSourceUrl	*string
	candidateType		*RTCStatsIceCandidateType
	ipAddress		*string
	portNumber		*int
	priority		*int
	transport		*string
}

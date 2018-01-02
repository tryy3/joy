package webrtc

type RTCIceCandidatePairStats struct {
	*RTCStats

	availableIncomingBitrate	*float32
	availableOutgoingBitrate	*float32
	bytesReceived			*uint64
	bytesSent			*uint64
	localCandidateId		*string
	nominated			*bool
	priority			*uint64
	readable			*bool
	remoteCandidateId		*string
	roundTripTime			*float32
	state				*RTCStatsIceCandidatePairState
	transportId			*string
	writable			*bool
}

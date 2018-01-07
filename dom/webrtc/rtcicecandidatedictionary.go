package webrtc

type RTCIceCandidateDictionary struct {
	foundation       *string
	ip               *string
	msMTurnSessionId *string
	port             *uint8
	priority         *uint
	protocol         *RTCIceProtocol
	relatedAddress   *string
	relatedPort      *uint8
	tcpType          *RTCIceTCPCandidateType
	kind             *RTCIceCandidateType
}

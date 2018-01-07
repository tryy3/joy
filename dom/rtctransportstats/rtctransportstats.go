package rtctransportstats

import "github.com/matthewmueller/joy/dom/webrtc"

type RTCTransportStats struct {
	*webrtc.RTCStats

	activeConnection        *bool
	bytesReceived           *uint64
	bytesSent               *uint64
	localCertificateId      *string
	remoteCertificateId     *string
	rtcpTransportStatsId    *string
	selectedCandidatePairId *string
}

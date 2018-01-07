package mstransportdiagnosticsstats

import (
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/webrtc"
)

type MSTransportDiagnosticsStats struct {
	*webrtc.RTCStats

	allocationTimeInMs     *uint
	baseAddress            *string
	baseInterface          *ms.MSNetworkInterfaceType
	iceRole                *webrtc.RTCIceRole
	iceWarningFlags        *ms.MSIceWarningFlags
	interfaces             *ms.MSNetworkInterfaceType
	localAddress           *string
	localAddrType          *ms.MSIceAddrType
	localInterface         *ms.MSNetworkInterfaceType
	localMR                *string
	localMRTCPPort         *uint8
	localSite              *string
	msRtcEngineVersion     *string
	networkName            *string
	numConsentReqReceived  *uint
	numConsentReqSent      *uint
	numConsentRespReceived *uint
	numConsentRespSent     *uint
	portRangeMax           *uint8
	portRangeMin           *uint8
	protocol               *webrtc.RTCIceProtocol
	remoteAddress          *string
	remoteAddrType         *ms.MSIceAddrType
	remoteMR               *string
	remoteMRTCPPort        *uint8
	remoteSite             *string
	rtpRtcpMux             *bool
	stunVer                *uint
}

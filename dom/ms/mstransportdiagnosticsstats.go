package ms

import "github.com/matthewmueller/joy/dom/webrtc"

type MSTransportDiagnosticsStats struct {
	*webrtc.RTCStats

	allocationTimeInMs	*uint
	baseAddress		*string
	baseInterface		*MSNetworkInterfaceType
	iceRole			*webrtc.RTCIceRole
	iceWarningFlags		*MSIceWarningFlags
	interfaces		*MSNetworkInterfaceType
	localAddress		*string
	localAddrType		*MSIceAddrType
	localInterface		*MSNetworkInterfaceType
	localMR			*string
	localMRTCPPort		*uint8
	localSite		*string
	msRtcEngineVersion	*string
	networkName		*string
	numConsentReqReceived	*uint
	numConsentReqSent	*uint
	numConsentRespReceived	*uint
	numConsentRespSent	*uint
	portRangeMax		*uint8
	portRangeMin		*uint8
	protocol		*webrtc.RTCIceProtocol
	remoteAddress		*string
	remoteAddrType		*MSIceAddrType
	remoteMR		*string
	remoteMRTCPPort		*uint8
	remoteSite		*string
	rtpRtcpMux		*bool
	stunVer			*uint
}

package ms

import "github.com/matthewmueller/joy/dom/webrtc"

type MSDescription struct {
	*webrtc.RTCStats

	connectivity		*MSConnectivity
	deviceDevName		*string
	localAddr		*MSIPAddressInfo
	networkconnectivity	*MSNetworkConnectivityInfo
	reflexiveLocalIPAddr	*MSIPAddressInfo
	remoteAddr		*MSIPAddressInfo
	transport		*webrtc.RTCIceProtocol
}

package msdescription

import (
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/webrtc"
)

type MSDescription struct {
	*webrtc.RTCStats

	connectivity         *ms.MSConnectivity
	deviceDevName        *string
	localAddr            *ms.MSIPAddressInfo
	networkconnectivity  *ms.MSNetworkConnectivityInfo
	reflexiveLocalIPAddr *ms.MSIPAddressInfo
	remoteAddr           *ms.MSIPAddressInfo
	transport            *webrtc.RTCIceProtocol
}

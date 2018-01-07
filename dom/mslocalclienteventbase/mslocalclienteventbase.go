package mslocalclienteventbase

import "github.com/matthewmueller/joy/dom/webrtc"

type MSLocalClientEventBase struct {
	*webrtc.RTCStats

	networkBandwidthLowEventRatio   *float32
	networkReceiveQualityEventRatio *float32
}

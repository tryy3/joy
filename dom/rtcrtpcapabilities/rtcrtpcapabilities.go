package rtcrtpcapabilities

import "github.com/matthewmueller/joy/dom/webrtc"

type RTCRtpCapabilities struct {
	codecs           *[]*webrtc.RTCRtpCodecCapability
	fecMechanisms    *[]string
	headerExtensions *[]*webrtc.RTCRtpHeaderExtension
}

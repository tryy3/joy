package rtcrtpparameters

import "github.com/matthewmueller/joy/dom/webrtc"

type RTCRtpParameters struct {
	codecs                *[]*webrtc.RTCRtpCodecParameters
	degradationPreference *webrtc.RTCDegradationPreference
	encodings             *[]*webrtc.RTCRtpEncodingParameters
	headerExtensions      *[]*webrtc.RTCRtpHeaderExtensionParameters
	muxId                 *string
	rtcp                  *webrtc.RTCRtcpParameters
}

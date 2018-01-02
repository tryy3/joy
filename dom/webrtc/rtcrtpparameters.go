package webrtc

type RTCRtpParameters struct {
	codecs			*[]*RTCRtpCodecParameters
	degradationPreference	*RTCDegradationPreference
	encodings		*[]*RTCRtpEncodingParameters
	headerExtensions	*[]*RTCRtpHeaderExtensionParameters
	muxId			*string
	rtcp			*RTCRtcpParameters
}

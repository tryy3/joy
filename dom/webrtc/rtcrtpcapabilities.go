package webrtc

type RTCRtpCapabilities struct {
	codecs			*[]*RTCRtpCodecCapability
	fecMechanisms		*[]string
	headerExtensions	*[]*RTCRtpHeaderExtension
}

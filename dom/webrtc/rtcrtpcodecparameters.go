package webrtc

type RTCRtpCodecParameters struct {
	clockRate    *uint
	maxptime     *uint
	name         *string
	numChannels  *uint
	parameters   *interface{}
	payloadType  *interface{}
	ptime        *uint
	rtcpFeedback *[]*RTCRtcpFeedback
}

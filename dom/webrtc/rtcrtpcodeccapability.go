package webrtc

type RTCRtpCodecCapability struct {
	clockRate             *uint
	kind                  *string
	maxptime              *uint
	maxSpatialLayers      *uint8
	maxTemporalLayers     *uint8
	name                  *string
	numChannels           *uint
	options               *interface{}
	parameters            *interface{}
	preferredPayloadType  *byte
	ptime                 *uint
	rtcpFeedback          *[]*RTCRtcpFeedback
	svcMultiStreamSupport *bool
}

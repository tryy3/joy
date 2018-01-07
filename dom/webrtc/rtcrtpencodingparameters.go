package webrtc

type RTCRtpEncodingParameters struct {
	active                *bool
	codecPayloadType      *byte
	dependencyEncodingIds *[]string
	encodingId            *string
	fec                   *RTCRtpFecParameters
	framerateScale        *float32
	maxBitrate            *float32
	maxFramerate          *uint
	minQuality            *float32
	priority              *float32
	resolutionScale       *float32
	rtx                   *RTCRtpRtxParameters
	ssrc                  *uint
	ssrcRange             *RTCSsrcRange
}

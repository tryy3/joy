package rtcmediastreamtrackstats

import "github.com/matthewmueller/joy/dom/webrtc"

type RTCMediaStreamTrackStats struct {
	*webrtc.RTCStats

	audioLevel                *float32
	echoReturnLoss            *float32
	echoReturnLossEnhancement *float32
	frameHeight               *uint
	framesCorrupted           *uint
	framesDecoded             *uint
	framesDropped             *uint
	framesPerSecond           *float32
	framesReceived            *uint
	framesSent                *uint
	frameWidth                *uint
	remoteSource              *bool
	ssrcIds                   *[]string
	trackIdentifier           *string
}

package mspayloadbase

import "github.com/matthewmueller/joy/dom/webrtc"

type MSPayloadBase struct {
	*webrtc.RTCStats

	payloadDescription *string
}

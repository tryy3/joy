package ms


import "github.com/matthewmueller/joy/dom/rtcstats"

type MSPayloadBase struct {
	*rtcstats.RTCStats

	payloadDescription *string
}

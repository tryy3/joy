package webrtc

import "github.com/matthewmueller/joy/dom/event"

type RTCStatsProvider interface {
	event.
		EventTarget

	GetStats() (r *RTCStatsReport)

	MsGetStats() (r *RTCStatsReport)
}

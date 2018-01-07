package webrtc

import "github.com/matthewmueller/joy/dom/ms"

type RTCStats struct {
	id        *string
	msType    *ms.MSStatsType
	timestamp *int
	kind      *RTCStatsType
}

package rtcstatsprovider

import (
	"github.com/matthewmueller/joy/dom/webrtc"
	"github.com/matthewmueller/joy/dom/window"
)

// RTCStatsProvider interface
// js:"RTCStatsProvider"
type RTCStatsProvider interface {
	window.EventTarget

	// GetStats
	// js:"getStats"
	// jsrewrite:"await $_.getStats()"
	GetStats() (w *webrtc.RTCStatsReport)

	// MsGetStats
	// js:"msGetStats"
	// jsrewrite:"await $_.msGetStats()"
	MsGetStats() (w *webrtc.RTCStatsReport)
}

package rtcicegatherer

import (
	"github.com/matthewmueller/joy/dom/rtcicegathererevent"
	"github.com/matthewmueller/joy/dom/rtcstatsprovider"
	"github.com/matthewmueller/joy/dom/webrtc"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ rtcstatsprovider.RTCStatsProvider = (*RTCIceGatherer)(nil)
var _ window.EventTarget = (*RTCIceGatherer)(nil)

// New fn
func New(options *webrtc.RTCIceGatherOptions) *RTCIceGatherer {
	macro.Rewrite("new RTCIceGatherer($1)", options)
	return &RTCIceGatherer{}
}

// RTCIceGatherer struct
// js:"RTCIceGatherer,omit"
type RTCIceGatherer struct {
}

// CreateAssociatedGatherer fn
// js:"createAssociatedGatherer"
func (*RTCIceGatherer) CreateAssociatedGatherer() (r *RTCIceGatherer) {
	macro.Rewrite("$_.createAssociatedGatherer()")
	return r
}

// GetLocalCandidates fn
// js:"getLocalCandidates"
func (*RTCIceGatherer) GetLocalCandidates() (w []*webrtc.RTCIceCandidateDictionary) {
	macro.Rewrite("$_.getLocalCandidates()")
	return w
}

// GetLocalParameters fn
// js:"getLocalParameters"
func (*RTCIceGatherer) GetLocalParameters() (w *webrtc.RTCIceParameters) {
	macro.Rewrite("$_.getLocalParameters()")
	return w
}

// GetStats fn
// js:"getStats"
func (*RTCIceGatherer) GetStats() (w *webrtc.RTCStatsReport) {
	macro.Rewrite("await $_.getStats()")
	return w
}

// MsGetStats fn
// js:"msGetStats"
func (*RTCIceGatherer) MsGetStats() (w *webrtc.RTCStatsReport) {
	macro.Rewrite("await $_.msGetStats()")
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCIceGatherer) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCIceGatherer) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCIceGatherer) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Component prop
// js:"component"
func (*RTCIceGatherer) Component() (component *webrtc.RTCIceComponent) {
	macro.Rewrite("$_.component")
	return component
}

// Onerror prop
// js:"onerror"
func (*RTCIceGatherer) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*RTCIceGatherer) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onlocalcandidate prop
// js:"onlocalcandidate"
func (*RTCIceGatherer) Onlocalcandidate() (onlocalcandidate func(*rtcicegathererevent.RTCIceGathererEvent)) {
	macro.Rewrite("$_.onlocalcandidate")
	return onlocalcandidate
}

// SetOnlocalcandidate prop
// js:"onlocalcandidate"
func (*RTCIceGatherer) SetOnlocalcandidate(onlocalcandidate func(*rtcicegathererevent.RTCIceGathererEvent)) {
	macro.Rewrite("$_.onlocalcandidate = $1", onlocalcandidate)
}

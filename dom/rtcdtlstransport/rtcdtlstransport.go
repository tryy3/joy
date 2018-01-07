package rtcdtlstransport

import (
	"github.com/matthewmueller/joy/dom/rtcdtlstransportstatechangedevent"
	"github.com/matthewmueller/joy/dom/rtcicetransport"
	"github.com/matthewmueller/joy/dom/rtcstatsprovider"
	"github.com/matthewmueller/joy/dom/webrtc"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ rtcstatsprovider.RTCStatsProvider = (*RTCDtlsTransport)(nil)
var _ window.EventTarget = (*RTCDtlsTransport)(nil)

// New fn
func New(transport *rtcicetransport.RTCIceTransport) *RTCDtlsTransport {
	macro.Rewrite("new RTCDtlsTransport($1)", transport)
	return &RTCDtlsTransport{}
}

// RTCDtlsTransport struct
// js:"RTCDtlsTransport,omit"
type RTCDtlsTransport struct {
}

// GetLocalParameters fn
// js:"getLocalParameters"
func (*RTCDtlsTransport) GetLocalParameters() (w *webrtc.RTCDtlsParameters) {
	macro.Rewrite("$_.getLocalParameters()")
	return w
}

// GetRemoteCertificates fn
// js:"getRemoteCertificates"
func (*RTCDtlsTransport) GetRemoteCertificates() (b [][]byte) {
	macro.Rewrite("$_.getRemoteCertificates()")
	return b
}

// GetRemoteParameters fn
// js:"getRemoteParameters"
func (*RTCDtlsTransport) GetRemoteParameters() (w *webrtc.RTCDtlsParameters) {
	macro.Rewrite("$_.getRemoteParameters()")
	return w
}

// Start fn
// js:"start"
func (*RTCDtlsTransport) Start(remoteParameters *webrtc.RTCDtlsParameters) {
	macro.Rewrite("$_.start($1)", remoteParameters)
}

// Stop fn
// js:"stop"
func (*RTCDtlsTransport) Stop() {
	macro.Rewrite("$_.stop()")
}

// GetStats fn
// js:"getStats"
func (*RTCDtlsTransport) GetStats() (w *webrtc.RTCStatsReport) {
	macro.Rewrite("await $_.getStats()")
	return w
}

// MsGetStats fn
// js:"msGetStats"
func (*RTCDtlsTransport) MsGetStats() (w *webrtc.RTCStatsReport) {
	macro.Rewrite("await $_.msGetStats()")
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCDtlsTransport) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCDtlsTransport) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCDtlsTransport) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Ondtlsstatechange prop
// js:"ondtlsstatechange"
func (*RTCDtlsTransport) Ondtlsstatechange() (ondtlsstatechange func(*rtcdtlstransportstatechangedevent.RTCDtlsTransportStateChangedEvent)) {
	macro.Rewrite("$_.ondtlsstatechange")
	return ondtlsstatechange
}

// SetOndtlsstatechange prop
// js:"ondtlsstatechange"
func (*RTCDtlsTransport) SetOndtlsstatechange(ondtlsstatechange func(*rtcdtlstransportstatechangedevent.RTCDtlsTransportStateChangedEvent)) {
	macro.Rewrite("$_.ondtlsstatechange = $1", ondtlsstatechange)
}

// Onerror prop
// js:"onerror"
func (*RTCDtlsTransport) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*RTCDtlsTransport) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// State prop
// js:"state"
func (*RTCDtlsTransport) State() (state *webrtc.RTCDtlsTransportState) {
	macro.Rewrite("$_.state")
	return state
}

// Transport prop
// js:"transport"
func (*RTCDtlsTransport) Transport() (transport *rtcicetransport.RTCIceTransport) {
	macro.Rewrite("$_.transport")
	return transport
}

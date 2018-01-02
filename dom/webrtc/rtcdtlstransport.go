package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ RTCStatsProvider = (*RTCDtlsTransport)(nil)
var _ event.EventTarget = (*RTCDtlsTransport)(nil)

func New(transport *RTCIceTransport) *RTCDtlsTransport {
	macro.Rewrite("new RTCDtlsTransport($1)", transport)
	return &RTCDtlsTransport{}
}

type RTCDtlsTransport struct {
}

func (*RTCDtlsTransport) GetLocalParameters() (r *RTCDtlsParameters) {
	macro.Rewrite("$_.getLocalParameters()")
	return r
}

func (*RTCDtlsTransport) GetRemoteCertificates() (b [][]byte) {
	macro.Rewrite("$_.getRemoteCertificates()")
	return b
}

func (*RTCDtlsTransport) GetRemoteParameters() (r *RTCDtlsParameters) {
	macro.Rewrite("$_.getRemoteParameters()")
	return r
}

func (*RTCDtlsTransport) Start(remoteParameters *RTCDtlsParameters) {
	macro.Rewrite("$_.start($1)", remoteParameters)
}

func (*RTCDtlsTransport) Stop() {
	macro.Rewrite("$_.stop()")
}

func (*RTCDtlsTransport) GetStats() (r *RTCStatsReport) {
	macro.Rewrite("await $_.getStats()")
	return r
}

func (*RTCDtlsTransport) MsGetStats() (r *RTCStatsReport) {
	macro.Rewrite("await $_.msGetStats()")
	return r
}

func (*RTCDtlsTransport) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCDtlsTransport) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*RTCDtlsTransport) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCDtlsTransport) Ondtlsstatechange() (ondtlsstatechange func(*RTCDtlsTransportStateChangedEvent)) {
	macro.Rewrite("$_.ondtlsstatechange")
	return ondtlsstatechange
}

func (*RTCDtlsTransport) SetOndtlsstatechange(ondtlsstatechange func(*RTCDtlsTransportStateChangedEvent)) {
	macro.Rewrite("$_.ondtlsstatechange = $1", ondtlsstatechange)
}

func (*RTCDtlsTransport) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*RTCDtlsTransport) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*RTCDtlsTransport) State() (state *RTCDtlsTransportState) {
	macro.Rewrite("$_.state")
	return state
}

func (*RTCDtlsTransport) Transport() (transport *RTCIceTransport) {
	macro.Rewrite("$_.transport")
	return transport
}

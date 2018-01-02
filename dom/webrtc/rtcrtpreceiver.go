package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/mediastreams"
)

var _ RTCStatsProvider = (*RTCRtpReceiver)(nil)
var _ event.EventTarget = (*RTCRtpReceiver)(nil)

func New(transport interface{}, kind string, rtcptransport *RTCDtlsTransport) *RTCRtpReceiver {
	macro.Rewrite("new RTCRtpReceiver($1, $2, $3)", transport, kind, rtcptransport)
	return &RTCRtpReceiver{}
}

type RTCRtpReceiver struct {
}

func (*RTCRtpReceiver) GetCapabilities(kind *string) (r *RTCRtpCapabilities) {
	macro.Rewrite("$_.getCapabilities($1)", kind)
	return r
}

func (*RTCRtpReceiver) GetContributingSources() (r []*RTCRtpContributingSource) {
	macro.Rewrite("$_.getContributingSources()")
	return r
}

func (*RTCRtpReceiver) Receive(parameters *RTCRtpParameters) {
	macro.Rewrite("$_.receive($1)", parameters)
}

func (*RTCRtpReceiver) RequestSendCSRC(csrc uint) {
	macro.Rewrite("$_.requestSendCSRC($1)", csrc)
}

func (*RTCRtpReceiver) SetTransport(transport interface{}, rtcpTransport *RTCDtlsTransport) {
	macro.Rewrite("$_.setTransport($1, $2)", transport, rtcpTransport)
}

func (*RTCRtpReceiver) Stop() {
	macro.Rewrite("$_.stop()")
}

func (*RTCRtpReceiver) GetStats() (r *RTCStatsReport) {
	macro.Rewrite("await $_.getStats()")
	return r
}

func (*RTCRtpReceiver) MsGetStats() (r *RTCStatsReport) {
	macro.Rewrite("await $_.msGetStats()")
	return r
}

func (*RTCRtpReceiver) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCRtpReceiver) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*RTCRtpReceiver) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCRtpReceiver) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*RTCRtpReceiver) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*RTCRtpReceiver) RtcpTransport() (rtcpTransport *RTCDtlsTransport) {
	macro.Rewrite("$_.rtcpTransport")
	return rtcpTransport
}

func (*RTCRtpReceiver) Track() (track *mediastreams.MediaStreamTrack) {
	macro.Rewrite("$_.track")
	return track
}

func (*RTCRtpReceiver) Transport() (transport interface{}) {
	macro.Rewrite("$_.transport")
	return transport
}

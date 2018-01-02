package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/mediastreams"
)

var _ RTCStatsProvider = (*RTCRtpSender)(nil)
var _ event.EventTarget = (*RTCRtpSender)(nil)

func New(track *mediastreams.MediaStreamTrack, transport interface{}, rtcptransport *RTCDtlsTransport) *RTCRtpSender {
	macro.Rewrite("new RTCRtpSender($1, $2, $3)", track, transport, rtcptransport)
	return &RTCRtpSender{}
}

type RTCRtpSender struct {
}

func (*RTCRtpSender) GetCapabilities(kind *string) (r *RTCRtpCapabilities) {
	macro.Rewrite("$_.getCapabilities($1)", kind)
	return r
}

func (*RTCRtpSender) Send(parameters *RTCRtpParameters) {
	macro.Rewrite("$_.send($1)", parameters)
}

func (*RTCRtpSender) SetTrack(track *mediastreams.MediaStreamTrack) {
	macro.Rewrite("$_.setTrack($1)", track)
}

func (*RTCRtpSender) SetTransport(transport interface{}, rtcpTransport *RTCDtlsTransport) {
	macro.Rewrite("$_.setTransport($1, $2)", transport, rtcpTransport)
}

func (*RTCRtpSender) Stop() {
	macro.Rewrite("$_.stop()")
}

func (*RTCRtpSender) GetStats() (r *RTCStatsReport) {
	macro.Rewrite("await $_.getStats()")
	return r
}

func (*RTCRtpSender) MsGetStats() (r *RTCStatsReport) {
	macro.Rewrite("await $_.msGetStats()")
	return r
}

func (*RTCRtpSender) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCRtpSender) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*RTCRtpSender) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCRtpSender) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*RTCRtpSender) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*RTCRtpSender) Onssrcconflict() (onssrcconflict func(*RTCSsrcConflictEvent)) {
	macro.Rewrite("$_.onssrcconflict")
	return onssrcconflict
}

func (*RTCRtpSender) SetOnssrcconflict(onssrcconflict func(*RTCSsrcConflictEvent)) {
	macro.Rewrite("$_.onssrcconflict = $1", onssrcconflict)
}

func (*RTCRtpSender) RtcpTransport() (rtcpTransport *RTCDtlsTransport) {
	macro.Rewrite("$_.rtcpTransport")
	return rtcpTransport
}

func (*RTCRtpSender) Track() (track *mediastreams.MediaStreamTrack) {
	macro.Rewrite("$_.track")
	return track
}

func (*RTCRtpSender) Transport() (transport interface{}) {
	macro.Rewrite("$_.transport")
	return transport
}

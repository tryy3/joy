package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*RTCDtmfSender)(nil)

func New(sender *RTCRtpSender) *RTCDtmfSender {
	macro.Rewrite("new RTCDtmfSender($1)", sender)
	return &RTCDtmfSender{}
}

type RTCDtmfSender struct {
}

func (*RTCDtmfSender) InsertDTMF(tones string, duration *int, interToneGap *int) {
	macro.Rewrite("$_.insertDTMF($1, $2, $3)", tones, duration, interToneGap)
}

func (*RTCDtmfSender) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCDtmfSender) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*RTCDtmfSender) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCDtmfSender) CanInsertDTMF() (canInsertDTMF bool) {
	macro.Rewrite("$_.canInsertDTMF")
	return canInsertDTMF
}

func (*RTCDtmfSender) Duration() (duration int) {
	macro.Rewrite("$_.duration")
	return duration
}

func (*RTCDtmfSender) InterToneGap() (interToneGap int) {
	macro.Rewrite("$_.interToneGap")
	return interToneGap
}

func (*RTCDtmfSender) Ontonechange() (ontonechange func(*RTCDTMFToneChangeEvent)) {
	macro.Rewrite("$_.ontonechange")
	return ontonechange
}

func (*RTCDtmfSender) SetOntonechange(ontonechange func(*RTCDTMFToneChangeEvent)) {
	macro.Rewrite("$_.ontonechange = $1", ontonechange)
}

func (*RTCDtmfSender) Sender() (sender *RTCRtpSender) {
	macro.Rewrite("$_.sender")
	return sender
}

func (*RTCDtmfSender) ToneBuffer() (toneBuffer string) {
	macro.Rewrite("$_.toneBuffer")
	return toneBuffer
}

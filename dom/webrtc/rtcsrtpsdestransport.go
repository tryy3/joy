package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*RTCSrtpSdesTransport)(nil)

func New(transport *RTCIceTransport, encryptparameters *RTCSrtpSdesParameters, decryptparameters *RTCSrtpSdesParameters) *RTCSrtpSdesTransport {
	macro.Rewrite("new RTCSrtpSdesTransport($1, $2, $3)", transport, encryptparameters, decryptparameters)
	return &RTCSrtpSdesTransport{}
}

type RTCSrtpSdesTransport struct {
}

func (*RTCSrtpSdesTransport) GetLocalParameters() (r []*RTCSrtpSdesParameters) {
	macro.Rewrite("$_.getLocalParameters()")
	return r
}

func (*RTCSrtpSdesTransport) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCSrtpSdesTransport) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*RTCSrtpSdesTransport) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCSrtpSdesTransport) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*RTCSrtpSdesTransport) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*RTCSrtpSdesTransport) Transport() (transport *RTCIceTransport) {
	macro.Rewrite("$_.transport")
	return transport
}

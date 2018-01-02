package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ RTCStatsProvider = (*RTCIceTransport)(nil)
var _ event.EventTarget = (*RTCIceTransport)(nil)

func New() *RTCIceTransport {
	macro.Rewrite("new RTCIceTransport()")
	return &RTCIceTransport{}
}

type RTCIceTransport struct {
}

func (*RTCIceTransport) AddRemoteCandidate(remoteCandidate interface{}) {
	macro.Rewrite("$_.addRemoteCandidate($1)", remoteCandidate)
}

func (*RTCIceTransport) CreateAssociatedTransport() (r *RTCIceTransport) {
	macro.Rewrite("$_.createAssociatedTransport()")
	return r
}

func (*RTCIceTransport) GetNominatedCandidatePair() (r *RTCIceCandidatePair) {
	macro.Rewrite("$_.getNominatedCandidatePair()")
	return r
}

func (*RTCIceTransport) GetRemoteCandidates() (r []*RTCIceCandidateDictionary) {
	macro.Rewrite("$_.getRemoteCandidates()")
	return r
}

func (*RTCIceTransport) GetRemoteParameters() (r *RTCIceParameters) {
	macro.Rewrite("$_.getRemoteParameters()")
	return r
}

func (*RTCIceTransport) SetRemoteCandidates(remoteCandidates []*RTCIceCandidateDictionary) {
	macro.Rewrite("$_.setRemoteCandidates($1)", remoteCandidates)
}

func (*RTCIceTransport) Start(gatherer *RTCIceGatherer, remoteParameters *RTCIceParameters, role *RTCIceRole) {
	macro.Rewrite("$_.start($1, $2, $3)", gatherer, remoteParameters, role)
}

func (*RTCIceTransport) Stop() {
	macro.Rewrite("$_.stop()")
}

func (*RTCIceTransport) GetStats() (r *RTCStatsReport) {
	macro.Rewrite("await $_.getStats()")
	return r
}

func (*RTCIceTransport) MsGetStats() (r *RTCStatsReport) {
	macro.Rewrite("await $_.msGetStats()")
	return r
}

func (*RTCIceTransport) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCIceTransport) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*RTCIceTransport) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCIceTransport) Component() (component *RTCIceComponent) {
	macro.Rewrite("$_.component")
	return component
}

func (*RTCIceTransport) IceGatherer() (iceGatherer *RTCIceGatherer) {
	macro.Rewrite("$_.iceGatherer")
	return iceGatherer
}

func (*RTCIceTransport) Oncandidatepairchange() (oncandidatepairchange func(*RTCIceCandidatePairChangedEvent)) {
	macro.Rewrite("$_.oncandidatepairchange")
	return oncandidatepairchange
}

func (*RTCIceTransport) SetOncandidatepairchange(oncandidatepairchange func(*RTCIceCandidatePairChangedEvent)) {
	macro.Rewrite("$_.oncandidatepairchange = $1", oncandidatepairchange)
}

func (*RTCIceTransport) Onicestatechange() (onicestatechange func(*RTCIceTransportStateChangedEvent)) {
	macro.Rewrite("$_.onicestatechange")
	return onicestatechange
}

func (*RTCIceTransport) SetOnicestatechange(onicestatechange func(*RTCIceTransportStateChangedEvent)) {
	macro.Rewrite("$_.onicestatechange = $1", onicestatechange)
}

func (*RTCIceTransport) Role() (role *RTCIceRole) {
	macro.Rewrite("$_.role")
	return role
}

func (*RTCIceTransport) State() (state *RTCIceTransportState) {
	macro.Rewrite("$_.state")
	return state
}

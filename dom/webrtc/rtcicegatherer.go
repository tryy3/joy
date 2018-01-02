package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ RTCStatsProvider = (*RTCIceGatherer)(nil)
var _ event.EventTarget = (*RTCIceGatherer)(nil)

func New(options *RTCIceGatherOptions) *RTCIceGatherer {
	macro.Rewrite("new RTCIceGatherer($1)", options)
	return &RTCIceGatherer{}
}

type RTCIceGatherer struct {
}

func (*RTCIceGatherer) CreateAssociatedGatherer() (r *RTCIceGatherer) {
	macro.Rewrite("$_.createAssociatedGatherer()")
	return r
}

func (*RTCIceGatherer) GetLocalCandidates() (r []*RTCIceCandidateDictionary) {
	macro.Rewrite("$_.getLocalCandidates()")
	return r
}

func (*RTCIceGatherer) GetLocalParameters() (r *RTCIceParameters) {
	macro.Rewrite("$_.getLocalParameters()")
	return r
}

func (*RTCIceGatherer) GetStats() (r *RTCStatsReport) {
	macro.Rewrite("await $_.getStats()")
	return r
}

func (*RTCIceGatherer) MsGetStats() (r *RTCStatsReport) {
	macro.Rewrite("await $_.msGetStats()")
	return r
}

func (*RTCIceGatherer) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCIceGatherer) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*RTCIceGatherer) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*RTCIceGatherer) Component() (component *RTCIceComponent) {
	macro.Rewrite("$_.component")
	return component
}

func (*RTCIceGatherer) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*RTCIceGatherer) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*RTCIceGatherer) Onlocalcandidate() (onlocalcandidate func(*RTCIceGathererEvent)) {
	macro.Rewrite("$_.onlocalcandidate")
	return onlocalcandidate
}

func (*RTCIceGatherer) SetOnlocalcandidate(onlocalcandidate func(*RTCIceGathererEvent)) {
	macro.Rewrite("$_.onlocalcandidate = $1", onlocalcandidate)
}

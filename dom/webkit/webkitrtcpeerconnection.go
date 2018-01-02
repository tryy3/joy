package webkit

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/webrtc"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/mediastreams"
)

var _ webrtc.RTCPeerConnection = (*WebkitRTCPeerConnection)(nil)
var _ event.EventTarget = (*WebkitRTCPeerConnection)(nil)

func New(configuration *webrtc.RTCConfiguration) *WebkitRTCPeerConnection {
	macro.Rewrite("new webkitRTCPeerConnection($1)", configuration)
	return &WebkitRTCPeerConnection{}
}

type WebkitRTCPeerConnection struct {
}

func (*WebkitRTCPeerConnection) AddIceCandidate(candidate *webrtc.RTCIceCandidate, successCallback *func(), failureCallback *func(err *dom.DOMError)) {
	macro.Rewrite("await $_.addIceCandidate($1, $2, $3)", candidate, successCallback, failureCallback)
}

func (*WebkitRTCPeerConnection) AddStream(stream *mediastreams.MediaStream) {
	macro.Rewrite("$_.addStream($1)", stream)
}

func (*WebkitRTCPeerConnection) Close() {
	macro.Rewrite("$_.close()")
}

func (*WebkitRTCPeerConnection) CreateAnswer(successCallback *func(sdp *webrtc.RTCSessionDescription), failureCallback *func(err *dom.DOMError)) (r *webrtc.RTCSessionDescription) {
	macro.Rewrite("await $_.createAnswer($1, $2)", successCallback, failureCallback)
	return r
}

func (*WebkitRTCPeerConnection) CreateOffer(successCallback *func(sdp *webrtc.RTCSessionDescription), failureCallback *func(err *dom.DOMError), options *webrtc.RTCOfferOptions) (r *webrtc.RTCSessionDescription) {
	macro.Rewrite("await $_.createOffer($1, $2, $3)", successCallback, failureCallback, options)
	return r
}

func (*WebkitRTCPeerConnection) GetConfiguration() (r *webrtc.RTCConfiguration) {
	macro.Rewrite("$_.getConfiguration()")
	return r
}

func (*WebkitRTCPeerConnection) GetLocalStreams() (w []*mediastreams.MediaStream) {
	macro.Rewrite("$_.getLocalStreams()")
	return w
}

func (*WebkitRTCPeerConnection) GetRemoteStreams() (w []*mediastreams.MediaStream) {
	macro.Rewrite("$_.getRemoteStreams()")
	return w
}

func (*WebkitRTCPeerConnection) GetStats(selector *mediastreams.MediaStreamTrack, successCallback *func(report *webrtc.RTCStatsReport), failureCallback *func(err *dom.DOMError)) (r *webrtc.RTCStatsReport) {
	macro.Rewrite("await $_.getStats($1, $2, $3)", selector, successCallback, failureCallback)
	return r
}

func (*WebkitRTCPeerConnection) GetStreamByID(streamId string) (w *mediastreams.MediaStream) {
	macro.Rewrite("$_.getStreamById($1)", streamId)
	return w
}

func (*WebkitRTCPeerConnection) RemoveStream(stream *mediastreams.MediaStream) {
	macro.Rewrite("$_.removeStream($1)", stream)
}

func (*WebkitRTCPeerConnection) SetLocalDescription(description *webrtc.RTCSessionDescription, successCallback *func(), failureCallback *func(err *dom.DOMError)) {
	macro.Rewrite("await $_.setLocalDescription($1, $2, $3)", description, successCallback, failureCallback)
}

func (*WebkitRTCPeerConnection) SetRemoteDescription(description *webrtc.RTCSessionDescription, successCallback *func(), failureCallback *func(err *dom.DOMError)) {
	macro.Rewrite("await $_.setRemoteDescription($1, $2, $3)", description, successCallback, failureCallback)
}

func (*WebkitRTCPeerConnection) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*WebkitRTCPeerConnection) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*WebkitRTCPeerConnection) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*WebkitRTCPeerConnection) CanTrickleIceCandidates() (canTrickleIceCandidates bool) {
	macro.Rewrite("$_.canTrickleIceCandidates")
	return canTrickleIceCandidates
}

func (*WebkitRTCPeerConnection) IceConnectionState() (iceConnectionState *webrtc.RTCIceConnectionState) {
	macro.Rewrite("$_.iceConnectionState")
	return iceConnectionState
}

func (*WebkitRTCPeerConnection) IceGatheringState() (iceGatheringState *webrtc.RTCIceGatheringState) {
	macro.Rewrite("$_.iceGatheringState")
	return iceGatheringState
}

func (*WebkitRTCPeerConnection) LocalDescription() (localDescription *webrtc.RTCSessionDescription) {
	macro.Rewrite("$_.localDescription")
	return localDescription
}

func (*WebkitRTCPeerConnection) Onaddstream() (onaddstream func(*mediastreams.MediaStreamEvent)) {
	macro.Rewrite("$_.onaddstream")
	return onaddstream
}

func (*WebkitRTCPeerConnection) SetOnaddstream(onaddstream func(*mediastreams.MediaStreamEvent)) {
	macro.Rewrite("$_.onaddstream = $1", onaddstream)
}

func (*WebkitRTCPeerConnection) Onicecandidate() (onicecandidate func(*webrtc.RTCPeerConnectionIceEvent)) {
	macro.Rewrite("$_.onicecandidate")
	return onicecandidate
}

func (*WebkitRTCPeerConnection) SetOnicecandidate(onicecandidate func(*webrtc.RTCPeerConnectionIceEvent)) {
	macro.Rewrite("$_.onicecandidate = $1", onicecandidate)
}

func (*WebkitRTCPeerConnection) Oniceconnectionstatechange() (oniceconnectionstatechange func(event.Event)) {
	macro.Rewrite("$_.oniceconnectionstatechange")
	return oniceconnectionstatechange
}

func (*WebkitRTCPeerConnection) SetOniceconnectionstatechange(oniceconnectionstatechange func(event.Event)) {
	macro.Rewrite("$_.oniceconnectionstatechange = $1", oniceconnectionstatechange)
}

func (*WebkitRTCPeerConnection) Onicegatheringstatechange() (onicegatheringstatechange func(event.Event)) {
	macro.Rewrite("$_.onicegatheringstatechange")
	return onicegatheringstatechange
}

func (*WebkitRTCPeerConnection) SetOnicegatheringstatechange(onicegatheringstatechange func(event.Event)) {
	macro.Rewrite("$_.onicegatheringstatechange = $1", onicegatheringstatechange)
}

func (*WebkitRTCPeerConnection) Onnegotiationneeded() (onnegotiationneeded func(event.Event)) {
	macro.Rewrite("$_.onnegotiationneeded")
	return onnegotiationneeded
}

func (*WebkitRTCPeerConnection) SetOnnegotiationneeded(onnegotiationneeded func(event.Event)) {
	macro.Rewrite("$_.onnegotiationneeded = $1", onnegotiationneeded)
}

func (*WebkitRTCPeerConnection) Onremovestream() (onremovestream func(*mediastreams.MediaStreamEvent)) {
	macro.Rewrite("$_.onremovestream")
	return onremovestream
}

func (*WebkitRTCPeerConnection) SetOnremovestream(onremovestream func(*mediastreams.MediaStreamEvent)) {
	macro.Rewrite("$_.onremovestream = $1", onremovestream)
}

func (*WebkitRTCPeerConnection) Onsignalingstatechange() (onsignalingstatechange func(event.Event)) {
	macro.Rewrite("$_.onsignalingstatechange")
	return onsignalingstatechange
}

func (*WebkitRTCPeerConnection) SetOnsignalingstatechange(onsignalingstatechange func(event.Event)) {
	macro.Rewrite("$_.onsignalingstatechange = $1", onsignalingstatechange)
}

func (*WebkitRTCPeerConnection) RemoteDescription() (remoteDescription *webrtc.RTCSessionDescription) {
	macro.Rewrite("$_.remoteDescription")
	return remoteDescription
}

func (*WebkitRTCPeerConnection) SignalingState() (signalingState *webrtc.RTCSignalingState) {
	macro.Rewrite("$_.signalingState")
	return signalingState
}

package webrtc

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/mediastreams"
)

type RTCPeerConnection interface {
	event.
		EventTarget

	AddIceCandidate(candidate *RTCIceCandidate, successCallback *func(), failureCallback *func(err *dom.DOMError))

	AddStream(stream *mediastreams.MediaStream)

	Close()

	CreateAnswer(successCallback *func(sdp *RTCSessionDescription), failureCallback *func(err *dom.DOMError)) (r *RTCSessionDescription)

	CreateOffer(successCallback *func(sdp *RTCSessionDescription), failureCallback *func(err *dom.DOMError), options *RTCOfferOptions) (r *RTCSessionDescription)

	GetConfiguration() (r *RTCConfiguration)

	GetLocalStreams() (w []*mediastreams.MediaStream)

	GetRemoteStreams() (w []*mediastreams.MediaStream)

	GetStats(selector *mediastreams.MediaStreamTrack, successCallback *func(report *RTCStatsReport), failureCallback *func(err *dom.DOMError)) (r *RTCStatsReport)

	GetStreamByID(streamId string) (w *mediastreams.MediaStream)

	RemoveStream(stream *mediastreams.MediaStream)

	SetLocalDescription(description *RTCSessionDescription, successCallback *func(), failureCallback *func(err *dom.DOMError))

	SetRemoteDescription(description *RTCSessionDescription, successCallback *func(), failureCallback *func(err *dom.DOMError))

	CanTrickleIceCandidates() (canTrickleIceCandidates bool)

	IceConnectionState() (iceConnectionState *RTCIceConnectionState)

	IceGatheringState() (iceGatheringState *RTCIceGatheringState)

	LocalDescription() (localDescription *RTCSessionDescription)

	Onaddstream() (onaddstream func(*mediastreams.MediaStreamEvent))

	SetOnaddstream(onaddstream func(*mediastreams.MediaStreamEvent))

	Onicecandidate() (onicecandidate func(*RTCPeerConnectionIceEvent))

	SetOnicecandidate(onicecandidate func(*RTCPeerConnectionIceEvent))

	Oniceconnectionstatechange() (oniceconnectionstatechange func(event.Event))

	SetOniceconnectionstatechange(oniceconnectionstatechange func(event.Event))

	Onicegatheringstatechange() (onicegatheringstatechange func(event.Event))

	SetOnicegatheringstatechange(onicegatheringstatechange func(event.Event))

	Onnegotiationneeded() (onnegotiationneeded func(event.Event))

	SetOnnegotiationneeded(onnegotiationneeded func(event.Event))

	Onremovestream() (onremovestream func(*mediastreams.MediaStreamEvent))

	SetOnremovestream(onremovestream func(*mediastreams.MediaStreamEvent))

	Onsignalingstatechange() (onsignalingstatechange func(event.Event))

	SetOnsignalingstatechange(onsignalingstatechange func(event.Event))

	RemoteDescription() (remoteDescription *RTCSessionDescription)

	SignalingState() (signalingState *RTCSignalingState)
}

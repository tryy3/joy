package mediastreams

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*MediaStream)(nil)

func NewMediaStream(streamortracks *interface{}) *MediaStream {
	macro.Rewrite("new MediaStream($1)", streamortracks)
	return &MediaStream{}
}

type MediaStream struct {
}

func (*MediaStream) AddTrack(track *MediaStreamTrack) {
	macro.Rewrite("$_.addTrack($1)", track)
}

func (*MediaStream) Clone() (m *MediaStream) {
	macro.Rewrite("$_.clone()")
	return m
}

func (*MediaStream) GetAudioTracks() (m []*MediaStreamTrack) {
	macro.Rewrite("$_.getAudioTracks()")
	return m
}

func (*MediaStream) GetTrackByID(trackId string) (m *MediaStreamTrack) {
	macro.Rewrite("$_.getTrackById($1)", trackId)
	return m
}

func (*MediaStream) GetTracks() (m []*MediaStreamTrack) {
	macro.Rewrite("$_.getTracks()")
	return m
}

func (*MediaStream) GetVideoTracks() (m []*MediaStreamTrack) {
	macro.Rewrite("$_.getVideoTracks()")
	return m
}

func (*MediaStream) RemoveTrack(track *MediaStreamTrack) {
	macro.Rewrite("$_.removeTrack($1)", track)
}

func (*MediaStream) Stop() {
	macro.Rewrite("$_.stop()")
}

func (*MediaStream) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaStream) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MediaStream) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaStream) Active() (active bool) {
	macro.Rewrite("$_.active")
	return active
}

func (*MediaStream) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*MediaStream) Onactive() (onactive func(event.Event)) {
	macro.Rewrite("$_.onactive")
	return onactive
}

func (*MediaStream) SetOnactive(onactive func(event.Event)) {
	macro.Rewrite("$_.onactive = $1", onactive)
}

func (*MediaStream) Onaddtrack() (onaddtrack func(*MediaStreamTrackEvent)) {
	macro.Rewrite("$_.onaddtrack")
	return onaddtrack
}

func (*MediaStream) SetOnaddtrack(onaddtrack func(*MediaStreamTrackEvent)) {
	macro.Rewrite("$_.onaddtrack = $1", onaddtrack)
}

func (*MediaStream) Oninactive() (oninactive func(event.Event)) {
	macro.Rewrite("$_.oninactive")
	return oninactive
}

func (*MediaStream) SetOninactive(oninactive func(event.Event)) {
	macro.Rewrite("$_.oninactive = $1", oninactive)
}

func (*MediaStream) Onremovetrack() (onremovetrack func(*MediaStreamTrackEvent)) {
	macro.Rewrite("$_.onremovetrack")
	return onremovetrack
}

func (*MediaStream) SetOnremovetrack(onremovetrack func(*MediaStreamTrackEvent)) {
	macro.Rewrite("$_.onremovetrack = $1", onremovetrack)
}

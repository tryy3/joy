package webvvt

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*TextTrack)(nil)

type TextTrack struct {
}

func (*TextTrack) AddCue(cue TextTrackCue) {
	macro.Rewrite("$_.addCue($1)", cue)
}

func (*TextTrack) RemoveCue(cue TextTrackCue) {
	macro.Rewrite("$_.removeCue($1)", cue)
}

func (*TextTrack) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*TextTrack) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*TextTrack) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*TextTrack) ActiveCues() (activeCues *TextTrackCueList) {
	macro.Rewrite("$_.activeCues")
	return activeCues
}

func (*TextTrack) Cues() (cues *TextTrackCueList) {
	macro.Rewrite("$_.cues")
	return cues
}

func (*TextTrack) InBandMetadataTrackDispatchType() (inBandMetadataTrackDispatchType string) {
	macro.Rewrite("$_.inBandMetadataTrackDispatchType")
	return inBandMetadataTrackDispatchType
}

func (*TextTrack) Kind() (kind string) {
	macro.Rewrite("$_.kind")
	return kind
}

func (*TextTrack) Label() (label string) {
	macro.Rewrite("$_.label")
	return label
}

func (*TextTrack) Language() (language string) {
	macro.Rewrite("$_.language")
	return language
}

func (*TextTrack) Mode() (mode interface{}) {
	macro.Rewrite("$_.mode")
	return mode
}

func (*TextTrack) SetMode(mode interface{}) {
	macro.Rewrite("$_.mode = $1", mode)
}

func (*TextTrack) Oncuechange() (oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange")
	return oncuechange
}

func (*TextTrack) SetOncuechange(oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange = $1", oncuechange)
}

func (*TextTrack) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*TextTrack) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*TextTrack) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*TextTrack) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*TextTrack) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

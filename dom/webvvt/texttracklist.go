package webvvt

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/utils"
)

var _ event.EventTarget = (*TextTrackList)(nil)

type TextTrackList struct {
}

func (*TextTrackList) Item(index uint) (t *TextTrack) {
	macro.Rewrite("$_.item($1)", index)
	return t
}

func (*TextTrackList) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*TextTrackList) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*TextTrackList) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*TextTrackList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*TextTrackList) Onaddtrack() (onaddtrack func(*utils.TrackEvent)) {
	macro.Rewrite("$_.onaddtrack")
	return onaddtrack
}

func (*TextTrackList) SetOnaddtrack(onaddtrack func(*utils.TrackEvent)) {
	macro.Rewrite("$_.onaddtrack = $1", onaddtrack)
}

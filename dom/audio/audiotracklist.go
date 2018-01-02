package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/utils"
)

var _ event.EventTarget = (*AudioTrackList)(nil)

type AudioTrackList struct {
}

func (*AudioTrackList) GetTrackByID(id string) (a *AudioTrack) {
	macro.Rewrite("$_.getTrackById($1)", id)
	return a
}

func (*AudioTrackList) Item(index uint) (a *AudioTrack) {
	macro.Rewrite("$_.item($1)", index)
	return a
}

func (*AudioTrackList) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*AudioTrackList) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*AudioTrackList) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*AudioTrackList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*AudioTrackList) Onaddtrack() (onaddtrack func(*utils.TrackEvent)) {
	macro.Rewrite("$_.onaddtrack")
	return onaddtrack
}

func (*AudioTrackList) SetOnaddtrack(onaddtrack func(*utils.TrackEvent)) {
	macro.Rewrite("$_.onaddtrack = $1", onaddtrack)
}

func (*AudioTrackList) Onchange() (onchange func(event.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

func (*AudioTrackList) SetOnchange(onchange func(event.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

func (*AudioTrackList) Onremovetrack() (onremovetrack func(*utils.TrackEvent)) {
	macro.Rewrite("$_.onremovetrack")
	return onremovetrack
}

func (*AudioTrackList) SetOnremovetrack(onremovetrack func(*utils.TrackEvent)) {
	macro.Rewrite("$_.onremovetrack = $1", onremovetrack)
}

package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/utils"
)

var _ event.EventTarget = (*VideoTrackList)(nil)

type VideoTrackList struct {
}

func (*VideoTrackList) GetTrackByID(id string) (v *VideoTrack) {
	macro.Rewrite("$_.getTrackById($1)", id)
	return v
}

func (*VideoTrackList) Item(index uint) (v *VideoTrack) {
	macro.Rewrite("$_.item($1)", index)
	return v
}

func (*VideoTrackList) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*VideoTrackList) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*VideoTrackList) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*VideoTrackList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*VideoTrackList) Onaddtrack() (onaddtrack func(*utils.TrackEvent)) {
	macro.Rewrite("$_.onaddtrack")
	return onaddtrack
}

func (*VideoTrackList) SetOnaddtrack(onaddtrack func(*utils.TrackEvent)) {
	macro.Rewrite("$_.onaddtrack = $1", onaddtrack)
}

func (*VideoTrackList) Onchange() (onchange func(event.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

func (*VideoTrackList) SetOnchange(onchange func(event.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

func (*VideoTrackList) Onremovetrack() (onremovetrack func(*utils.TrackEvent)) {
	macro.Rewrite("$_.onremovetrack")
	return onremovetrack
}

func (*VideoTrackList) SetOnremovetrack(onremovetrack func(*utils.TrackEvent)) {
	macro.Rewrite("$_.onremovetrack = $1", onremovetrack)
}

func (*VideoTrackList) SelectedIndex() (selectedIndex int) {
	macro.Rewrite("$_.selectedIndex")
	return selectedIndex
}

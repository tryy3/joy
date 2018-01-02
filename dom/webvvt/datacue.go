package webvvt

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/document"
)

var _ TextTrackCue = (*DataCue)(nil)
var _ event.EventTarget = (*DataCue)(nil)

type DataCue struct {
}

func (*DataCue) GetCueAsHTML() (w *document.DocumentFragment) {
	macro.Rewrite("$_.getCueAsHTML()")
	return w
}

func (*DataCue) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*DataCue) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*DataCue) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*DataCue) Data() (data []byte) {
	macro.Rewrite("$_.data")
	return data
}

func (*DataCue) SetData(data []byte) {
	macro.Rewrite("$_.data = $1", data)
}

func (*DataCue) EndTime() (endTime float32) {
	macro.Rewrite("$_.endTime")
	return endTime
}

func (*DataCue) SetEndTime(endTime float32) {
	macro.Rewrite("$_.endTime = $1", endTime)
}

func (*DataCue) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*DataCue) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

func (*DataCue) Onenter() (onenter func(event.Event)) {
	macro.Rewrite("$_.onenter")
	return onenter
}

func (*DataCue) SetOnenter(onenter func(event.Event)) {
	macro.Rewrite("$_.onenter = $1", onenter)
}

func (*DataCue) Onexit() (onexit func(event.Event)) {
	macro.Rewrite("$_.onexit")
	return onexit
}

func (*DataCue) SetOnexit(onexit func(event.Event)) {
	macro.Rewrite("$_.onexit = $1", onexit)
}

func (*DataCue) PauseOnExit() (pauseOnExit bool) {
	macro.Rewrite("$_.pauseOnExit")
	return pauseOnExit
}

func (*DataCue) SetPauseOnExit(pauseOnExit bool) {
	macro.Rewrite("$_.pauseOnExit = $1", pauseOnExit)
}

func (*DataCue) StartTime() (startTime float32) {
	macro.Rewrite("$_.startTime")
	return startTime
}

func (*DataCue) SetStartTime(startTime float32) {
	macro.Rewrite("$_.startTime = $1", startTime)
}

func (*DataCue) Text() (text string) {
	macro.Rewrite("$_.text")
	return text
}

func (*DataCue) SetText(text string) {
	macro.Rewrite("$_.text = $1", text)
}

func (*DataCue) Track() (track *TextTrack) {
	macro.Rewrite("$_.track")
	return track
}

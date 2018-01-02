package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/audio"
)

var _ event.EventTarget = (*SourceBufferList)(nil)

type SourceBufferList struct {
}

func (*SourceBufferList) Item(index uint) (a *audio.SourceBuffer) {
	macro.Rewrite("$_.item($1)", index)
	return a
}

func (*SourceBufferList) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SourceBufferList) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*SourceBufferList) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SourceBufferList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

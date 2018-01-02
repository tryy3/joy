package mediasource

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/audio"
	"github.com/matthewmueller/joy/dom/utils"
)

var _ event.EventTarget = (*MediaSource)(nil)

func New() *MediaSource {
	macro.Rewrite("new MediaSource()")
	return &MediaSource{}
}

type MediaSource struct {
}

func (*MediaSource) AddSourceBuffer(kind string) (a *audio.SourceBuffer) {
	macro.Rewrite("$_.addSourceBuffer($1)", kind)
	return a
}

func (*MediaSource) EndOfStream(err *string) {
	macro.Rewrite("$_.endOfStream($1)", err)
}

func (*MediaSource) IsTypeSupported(kind string) (b bool) {
	macro.Rewrite("$_.isTypeSupported($1)", kind)
	return b
}

func (*MediaSource) RemoveSourceBuffer(sourceBuffer *audio.SourceBuffer) {
	macro.Rewrite("$_.removeSourceBuffer($1)", sourceBuffer)
}

func (*MediaSource) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaSource) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MediaSource) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaSource) ActiveSourceBuffers() (activeSourceBuffers *utils.SourceBufferList) {
	macro.Rewrite("$_.activeSourceBuffers")
	return activeSourceBuffers
}

func (*MediaSource) Duration() (duration float32) {
	macro.Rewrite("$_.duration")
	return duration
}

func (*MediaSource) SetDuration(duration float32) {
	macro.Rewrite("$_.duration = $1", duration)
}

func (*MediaSource) ReadyState() (readyState string) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*MediaSource) SourceBuffers() (sourceBuffers *utils.SourceBufferList) {
	macro.Rewrite("$_.sourceBuffers")
	return sourceBuffers
}

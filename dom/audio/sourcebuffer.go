package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/mediasource"
)

var _ event.EventTarget = (*SourceBuffer)(nil)

type SourceBuffer struct {
}

func (*SourceBuffer) Abort() {
	macro.Rewrite("$_.abort()")
}

func (*SourceBuffer) AppendBuffer(data interface{}) {
	macro.Rewrite("$_.appendBuffer($1)", data)
}

func (*SourceBuffer) AppendStream(stream *ms.MSStream, maxSize *uint64) {
	macro.Rewrite("$_.appendStream($1, $2)", stream, maxSize)
}

func (*SourceBuffer) Remove(start float32, end float32) {
	macro.Rewrite("$_.remove($1, $2)", start, end)
}

func (*SourceBuffer) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SourceBuffer) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*SourceBuffer) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SourceBuffer) AppendWindowEnd() (appendWindowEnd float32) {
	macro.Rewrite("$_.appendWindowEnd")
	return appendWindowEnd
}

func (*SourceBuffer) SetAppendWindowEnd(appendWindowEnd float32) {
	macro.Rewrite("$_.appendWindowEnd = $1", appendWindowEnd)
}

func (*SourceBuffer) AppendWindowStart() (appendWindowStart float32) {
	macro.Rewrite("$_.appendWindowStart")
	return appendWindowStart
}

func (*SourceBuffer) SetAppendWindowStart(appendWindowStart float32) {
	macro.Rewrite("$_.appendWindowStart = $1", appendWindowStart)
}

func (*SourceBuffer) AudioTracks() (audioTracks *AudioTrackList) {
	macro.Rewrite("$_.audioTracks")
	return audioTracks
}

func (*SourceBuffer) Buffered() (buffered *dom.TimeRanges) {
	macro.Rewrite("$_.buffered")
	return buffered
}

func (*SourceBuffer) Mode() (mode *mediasource.AppendMode) {
	macro.Rewrite("$_.mode")
	return mode
}

func (*SourceBuffer) SetMode(mode *mediasource.AppendMode) {
	macro.Rewrite("$_.mode = $1", mode)
}

func (*SourceBuffer) TimestampOffset() (timestampOffset float32) {
	macro.Rewrite("$_.timestampOffset")
	return timestampOffset
}

func (*SourceBuffer) SetTimestampOffset(timestampOffset float32) {
	macro.Rewrite("$_.timestampOffset = $1", timestampOffset)
}

func (*SourceBuffer) Updating() (updating bool) {
	macro.Rewrite("$_.updating")
	return updating
}

func (*SourceBuffer) VideoTracks() (videoTracks *VideoTrackList) {
	macro.Rewrite("$_.videoTracks")
	return videoTracks
}

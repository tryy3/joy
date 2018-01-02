package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*AudioBufferSourceNode)(nil)
var _ event.EventTarget = (*AudioBufferSourceNode)(nil)

type AudioBufferSourceNode struct {
}

func (*AudioBufferSourceNode) Start(when *float32, offset *float32, duration *float32) {
	macro.Rewrite("$_.start($1, $2, $3)", when, offset, duration)
}

func (*AudioBufferSourceNode) Stop(when *float32) {
	macro.Rewrite("$_.stop($1)", when)
}

func (*AudioBufferSourceNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*AudioBufferSourceNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*AudioBufferSourceNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*AudioBufferSourceNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*AudioBufferSourceNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*AudioBufferSourceNode) Buffer() (buffer *AudioBuffer) {
	macro.Rewrite("$_.buffer")
	return buffer
}

func (*AudioBufferSourceNode) SetBuffer(buffer *AudioBuffer) {
	macro.Rewrite("$_.buffer = $1", buffer)
}

func (*AudioBufferSourceNode) Detune() (detune *AudioParam) {
	macro.Rewrite("$_.detune")
	return detune
}

func (*AudioBufferSourceNode) Loop() (loop bool) {
	macro.Rewrite("$_.loop")
	return loop
}

func (*AudioBufferSourceNode) SetLoop(loop bool) {
	macro.Rewrite("$_.loop = $1", loop)
}

func (*AudioBufferSourceNode) LoopEnd() (loopEnd float32) {
	macro.Rewrite("$_.loopEnd")
	return loopEnd
}

func (*AudioBufferSourceNode) SetLoopEnd(loopEnd float32) {
	macro.Rewrite("$_.loopEnd = $1", loopEnd)
}

func (*AudioBufferSourceNode) LoopStart() (loopStart float32) {
	macro.Rewrite("$_.loopStart")
	return loopStart
}

func (*AudioBufferSourceNode) SetLoopStart(loopStart float32) {
	macro.Rewrite("$_.loopStart = $1", loopStart)
}

func (*AudioBufferSourceNode) Onended() (onended func(event.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*AudioBufferSourceNode) SetOnended(onended func(event.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*AudioBufferSourceNode) PlaybackRate() (playbackRate *AudioParam) {
	macro.Rewrite("$_.playbackRate")
	return playbackRate
}

func (*AudioBufferSourceNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*AudioBufferSourceNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*AudioBufferSourceNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*AudioBufferSourceNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*AudioBufferSourceNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*AudioBufferSourceNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*AudioBufferSourceNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*AudioBufferSourceNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*AudioBufferSourceNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

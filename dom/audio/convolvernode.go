package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*ConvolverNode)(nil)
var _ event.EventTarget = (*ConvolverNode)(nil)

type ConvolverNode struct {
}

func (*ConvolverNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*ConvolverNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*ConvolverNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ConvolverNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*ConvolverNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ConvolverNode) Buffer() (buffer *AudioBuffer) {
	macro.Rewrite("$_.buffer")
	return buffer
}

func (*ConvolverNode) SetBuffer(buffer *AudioBuffer) {
	macro.Rewrite("$_.buffer = $1", buffer)
}

func (*ConvolverNode) Normalize() (normalize bool) {
	macro.Rewrite("$_.normalize")
	return normalize
}

func (*ConvolverNode) SetNormalize(normalize bool) {
	macro.Rewrite("$_.normalize = $1", normalize)
}

func (*ConvolverNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*ConvolverNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*ConvolverNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*ConvolverNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*ConvolverNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*ConvolverNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*ConvolverNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*ConvolverNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*ConvolverNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

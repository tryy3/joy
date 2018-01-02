package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*MediaStreamAudioSourceNode)(nil)
var _ event.EventTarget = (*MediaStreamAudioSourceNode)(nil)

type MediaStreamAudioSourceNode struct {
}

func (*MediaStreamAudioSourceNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*MediaStreamAudioSourceNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*MediaStreamAudioSourceNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaStreamAudioSourceNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MediaStreamAudioSourceNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaStreamAudioSourceNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*MediaStreamAudioSourceNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*MediaStreamAudioSourceNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*MediaStreamAudioSourceNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*MediaStreamAudioSourceNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*MediaStreamAudioSourceNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*MediaStreamAudioSourceNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*MediaStreamAudioSourceNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*MediaStreamAudioSourceNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

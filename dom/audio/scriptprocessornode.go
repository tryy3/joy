package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*ScriptProcessorNode)(nil)
var _ event.EventTarget = (*ScriptProcessorNode)(nil)

type ScriptProcessorNode struct {
}

func (*ScriptProcessorNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*ScriptProcessorNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*ScriptProcessorNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ScriptProcessorNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*ScriptProcessorNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ScriptProcessorNode) BufferSize() (bufferSize int) {
	macro.Rewrite("$_.bufferSize")
	return bufferSize
}

func (*ScriptProcessorNode) Onaudioprocess() (onaudioprocess func(*AudioProcessingEvent)) {
	macro.Rewrite("$_.onaudioprocess")
	return onaudioprocess
}

func (*ScriptProcessorNode) SetOnaudioprocess(onaudioprocess func(*AudioProcessingEvent)) {
	macro.Rewrite("$_.onaudioprocess = $1", onaudioprocess)
}

func (*ScriptProcessorNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*ScriptProcessorNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*ScriptProcessorNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*ScriptProcessorNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*ScriptProcessorNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*ScriptProcessorNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*ScriptProcessorNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*ScriptProcessorNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*ScriptProcessorNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

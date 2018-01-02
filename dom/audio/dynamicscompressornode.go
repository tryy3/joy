package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*DynamicsCompressorNode)(nil)
var _ event.EventTarget = (*DynamicsCompressorNode)(nil)

type DynamicsCompressorNode struct {
}

func (*DynamicsCompressorNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*DynamicsCompressorNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*DynamicsCompressorNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*DynamicsCompressorNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*DynamicsCompressorNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*DynamicsCompressorNode) Attack() (attack *AudioParam) {
	macro.Rewrite("$_.attack")
	return attack
}

func (*DynamicsCompressorNode) Knee() (knee *AudioParam) {
	macro.Rewrite("$_.knee")
	return knee
}

func (*DynamicsCompressorNode) Ratio() (ratio *AudioParam) {
	macro.Rewrite("$_.ratio")
	return ratio
}

func (*DynamicsCompressorNode) Reduction() (reduction float32) {
	macro.Rewrite("$_.reduction")
	return reduction
}

func (*DynamicsCompressorNode) Release() (release *AudioParam) {
	macro.Rewrite("$_.release")
	return release
}

func (*DynamicsCompressorNode) Threshold() (threshold *AudioParam) {
	macro.Rewrite("$_.threshold")
	return threshold
}

func (*DynamicsCompressorNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*DynamicsCompressorNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*DynamicsCompressorNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*DynamicsCompressorNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*DynamicsCompressorNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*DynamicsCompressorNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*DynamicsCompressorNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*DynamicsCompressorNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*DynamicsCompressorNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

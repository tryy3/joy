package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*BiquadFilterNode)(nil)
var _ event.EventTarget = (*BiquadFilterNode)(nil)

type BiquadFilterNode struct {
}

func (*BiquadFilterNode) GetFrequencyResponse(frequencyHz []float32, magResponse []float32, phaseResponse []float32) {
	macro.Rewrite("$_.getFrequencyResponse($1, $2, $3)", frequencyHz, magResponse, phaseResponse)
}

func (*BiquadFilterNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*BiquadFilterNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*BiquadFilterNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*BiquadFilterNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*BiquadFilterNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*BiquadFilterNode) Detune() (detune *AudioParam) {
	macro.Rewrite("$_.detune")
	return detune
}

func (*BiquadFilterNode) Frequency() (frequency *AudioParam) {
	macro.Rewrite("$_.frequency")
	return frequency
}

func (*BiquadFilterNode) Gain() (gain *AudioParam) {
	macro.Rewrite("$_.gain")
	return gain
}

func (*BiquadFilterNode) Q() (Q *AudioParam) {
	macro.Rewrite("$_.Q")
	return Q
}

func (*BiquadFilterNode) Type() (kind *BiquadFilterType) {
	macro.Rewrite("$_.type")
	return kind
}

func (*BiquadFilterNode) SetType(kind *BiquadFilterType) {
	macro.Rewrite("$_.type = $1", kind)
}

func (*BiquadFilterNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*BiquadFilterNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*BiquadFilterNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*BiquadFilterNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*BiquadFilterNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*BiquadFilterNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*BiquadFilterNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*BiquadFilterNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*BiquadFilterNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

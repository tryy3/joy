package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*WaveShaperNode)(nil)
var _ event.EventTarget = (*WaveShaperNode)(nil)

type WaveShaperNode struct {
}

func (*WaveShaperNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*WaveShaperNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*WaveShaperNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*WaveShaperNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*WaveShaperNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*WaveShaperNode) Curve() (curve []float32) {
	macro.Rewrite("$_.curve")
	return curve
}

func (*WaveShaperNode) SetCurve(curve []float32) {
	macro.Rewrite("$_.curve = $1", curve)
}

func (*WaveShaperNode) Oversample() (oversample *OverSampleType) {
	macro.Rewrite("$_.oversample")
	return oversample
}

func (*WaveShaperNode) SetOversample(oversample *OverSampleType) {
	macro.Rewrite("$_.oversample = $1", oversample)
}

func (*WaveShaperNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*WaveShaperNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*WaveShaperNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*WaveShaperNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*WaveShaperNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*WaveShaperNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*WaveShaperNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*WaveShaperNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*WaveShaperNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

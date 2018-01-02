package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*AnalyserNode)(nil)
var _ event.EventTarget = (*AnalyserNode)(nil)

type AnalyserNode struct {
}

func (*AnalyserNode) GetByteFrequencyData(array []uint8) {
	macro.Rewrite("$_.getByteFrequencyData($1)", array)
}

func (*AnalyserNode) GetByteTimeDomainData(array []uint8) {
	macro.Rewrite("$_.getByteTimeDomainData($1)", array)
}

func (*AnalyserNode) GetFloatFrequencyData(array []float32) {
	macro.Rewrite("$_.getFloatFrequencyData($1)", array)
}

func (*AnalyserNode) GetFloatTimeDomainData(array []float32) {
	macro.Rewrite("$_.getFloatTimeDomainData($1)", array)
}

func (*AnalyserNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*AnalyserNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*AnalyserNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*AnalyserNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*AnalyserNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*AnalyserNode) FftSize() (fftSize uint) {
	macro.Rewrite("$_.fftSize")
	return fftSize
}

func (*AnalyserNode) SetFftSize(fftSize uint) {
	macro.Rewrite("$_.fftSize = $1", fftSize)
}

func (*AnalyserNode) FrequencyBinCount() (frequencyBinCount uint) {
	macro.Rewrite("$_.frequencyBinCount")
	return frequencyBinCount
}

func (*AnalyserNode) MaxDecibels() (maxDecibels float32) {
	macro.Rewrite("$_.maxDecibels")
	return maxDecibels
}

func (*AnalyserNode) SetMaxDecibels(maxDecibels float32) {
	macro.Rewrite("$_.maxDecibels = $1", maxDecibels)
}

func (*AnalyserNode) MinDecibels() (minDecibels float32) {
	macro.Rewrite("$_.minDecibels")
	return minDecibels
}

func (*AnalyserNode) SetMinDecibels(minDecibels float32) {
	macro.Rewrite("$_.minDecibels = $1", minDecibels)
}

func (*AnalyserNode) SmoothingTimeConstant() (smoothingTimeConstant float32) {
	macro.Rewrite("$_.smoothingTimeConstant")
	return smoothingTimeConstant
}

func (*AnalyserNode) SetSmoothingTimeConstant(smoothingTimeConstant float32) {
	macro.Rewrite("$_.smoothingTimeConstant = $1", smoothingTimeConstant)
}

func (*AnalyserNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*AnalyserNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*AnalyserNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*AnalyserNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*AnalyserNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*AnalyserNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*AnalyserNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*AnalyserNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*AnalyserNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

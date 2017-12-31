package audio


import (
	"github.com/matthewmueller/joy/dom/channelcountmode"
	"github.com/matthewmueller/joy/dom/channelinterpretation"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ AudioNode = (*AnalyserNode)(nil)
var _ window.EventTarget = (*AnalyserNode)(nil)

// AnalyserNode struct
// js:"AnalyserNode,omit"
type AnalyserNode struct {
}

// GetByteFrequencyData fn
// js:"getByteFrequencyData"
func (*AnalyserNode) GetByteFrequencyData(array []uint8) {
	macro.Rewrite("$_.getByteFrequencyData($1)", array)
}

// GetByteTimeDomainData fn
// js:"getByteTimeDomainData"
func (*AnalyserNode) GetByteTimeDomainData(array []uint8) {
	macro.Rewrite("$_.getByteTimeDomainData($1)", array)
}

// GetFloatFrequencyData fn
// js:"getFloatFrequencyData"
func (*AnalyserNode) GetFloatFrequencyData(array []float32) {
	macro.Rewrite("$_.getFloatFrequencyData($1)", array)
}

// GetFloatTimeDomainData fn
// js:"getFloatTimeDomainData"
func (*AnalyserNode) GetFloatTimeDomainData(array []float32) {
	macro.Rewrite("$_.getFloatTimeDomainData($1)", array)
}

// Connect fn
// js:"connect"
func (*AnalyserNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*AnalyserNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*AnalyserNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*AnalyserNode) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*AnalyserNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// FftSize prop
// js:"fftSize"
func (*AnalyserNode) FftSize() (fftSize uint) {
	macro.Rewrite("$_.fftSize")
	return fftSize
}

// SetFftSize prop
// js:"fftSize"
func (*AnalyserNode) SetFftSize(fftSize uint) {
	macro.Rewrite("$_.fftSize = $1", fftSize)
}

// FrequencyBinCount prop
// js:"frequencyBinCount"
func (*AnalyserNode) FrequencyBinCount() (frequencyBinCount uint) {
	macro.Rewrite("$_.frequencyBinCount")
	return frequencyBinCount
}

// MaxDecibels prop
// js:"maxDecibels"
func (*AnalyserNode) MaxDecibels() (maxDecibels float32) {
	macro.Rewrite("$_.maxDecibels")
	return maxDecibels
}

// SetMaxDecibels prop
// js:"maxDecibels"
func (*AnalyserNode) SetMaxDecibels(maxDecibels float32) {
	macro.Rewrite("$_.maxDecibels = $1", maxDecibels)
}

// MinDecibels prop
// js:"minDecibels"
func (*AnalyserNode) MinDecibels() (minDecibels float32) {
	macro.Rewrite("$_.minDecibels")
	return minDecibels
}

// SetMinDecibels prop
// js:"minDecibels"
func (*AnalyserNode) SetMinDecibels(minDecibels float32) {
	macro.Rewrite("$_.minDecibels = $1", minDecibels)
}

// SmoothingTimeConstant prop
// js:"smoothingTimeConstant"
func (*AnalyserNode) SmoothingTimeConstant() (smoothingTimeConstant float32) {
	macro.Rewrite("$_.smoothingTimeConstant")
	return smoothingTimeConstant
}

// SetSmoothingTimeConstant prop
// js:"smoothingTimeConstant"
func (*AnalyserNode) SetSmoothingTimeConstant(smoothingTimeConstant float32) {
	macro.Rewrite("$_.smoothingTimeConstant = $1", smoothingTimeConstant)
}

// ChannelCount prop
// js:"channelCount"
func (*AnalyserNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*AnalyserNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*AnalyserNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*AnalyserNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*AnalyserNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*AnalyserNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*AnalyserNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*AnalyserNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*AnalyserNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*OscillatorNode)(nil)
var _ event.EventTarget = (*OscillatorNode)(nil)

type OscillatorNode struct {
}

func (*OscillatorNode) SetPeriodicWave(periodicWave *PeriodicWave) {
	macro.Rewrite("$_.setPeriodicWave($1)", periodicWave)
}

func (*OscillatorNode) Start(when *float32) {
	macro.Rewrite("$_.start($1)", when)
}

func (*OscillatorNode) Stop(when *float32) {
	macro.Rewrite("$_.stop($1)", when)
}

func (*OscillatorNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*OscillatorNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*OscillatorNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*OscillatorNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*OscillatorNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*OscillatorNode) Detune() (detune *AudioParam) {
	macro.Rewrite("$_.detune")
	return detune
}

func (*OscillatorNode) Frequency() (frequency *AudioParam) {
	macro.Rewrite("$_.frequency")
	return frequency
}

func (*OscillatorNode) Onended() (onended func(event.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*OscillatorNode) SetOnended(onended func(event.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*OscillatorNode) Type() (kind *OscillatorType) {
	macro.Rewrite("$_.type")
	return kind
}

func (*OscillatorNode) SetType(kind *OscillatorType) {
	macro.Rewrite("$_.type = $1", kind)
}

func (*OscillatorNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*OscillatorNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*OscillatorNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*OscillatorNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*OscillatorNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*OscillatorNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*OscillatorNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*OscillatorNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*OscillatorNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

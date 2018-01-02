package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*StereoPannerNode)(nil)
var _ event.EventTarget = (*StereoPannerNode)(nil)

type StereoPannerNode struct {
}

func (*StereoPannerNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*StereoPannerNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*StereoPannerNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*StereoPannerNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*StereoPannerNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*StereoPannerNode) Pan() (pan *AudioParam) {
	macro.Rewrite("$_.pan")
	return pan
}

func (*StereoPannerNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*StereoPannerNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*StereoPannerNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*StereoPannerNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*StereoPannerNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*StereoPannerNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*StereoPannerNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*StereoPannerNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*StereoPannerNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

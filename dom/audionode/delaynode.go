package audionode

import (
	"github.com/matthewmueller/joy/dom/audio"
	"github.com/matthewmueller/joy/dom/audioparam"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ AudioNode = (*DelayNode)(nil)
var _ window.EventTarget = (*DelayNode)(nil)

// DelayNode struct
// js:"DelayNode,omit"
type DelayNode struct {
}

// Connect fn
// js:"connect"
func (*DelayNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*DelayNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*DelayNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*DelayNode) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*DelayNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DelayTime prop
// js:"delayTime"
func (*DelayNode) DelayTime() (delayTime *audioparam.AudioParam) {
	macro.Rewrite("$_.delayTime")
	return delayTime
}

// ChannelCount prop
// js:"channelCount"
func (*DelayNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*DelayNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*DelayNode) ChannelCountMode() (channelCountMode *audio.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*DelayNode) SetChannelCountMode(channelCountMode *audio.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*DelayNode) ChannelInterpretation() (channelInterpretation *audio.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*DelayNode) SetChannelInterpretation(channelInterpretation *audio.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*DelayNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*DelayNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*DelayNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

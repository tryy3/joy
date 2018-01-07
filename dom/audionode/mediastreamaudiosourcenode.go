package audionode

import (
	"github.com/matthewmueller/joy/dom/audio"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ AudioNode = (*MediaStreamAudioSourceNode)(nil)
var _ window.EventTarget = (*MediaStreamAudioSourceNode)(nil)

// MediaStreamAudioSourceNode struct
// js:"MediaStreamAudioSourceNode,omit"
type MediaStreamAudioSourceNode struct {
}

// Connect fn
// js:"connect"
func (*MediaStreamAudioSourceNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*MediaStreamAudioSourceNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*MediaStreamAudioSourceNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MediaStreamAudioSourceNode) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MediaStreamAudioSourceNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ChannelCount prop
// js:"channelCount"
func (*MediaStreamAudioSourceNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*MediaStreamAudioSourceNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*MediaStreamAudioSourceNode) ChannelCountMode() (channelCountMode *audio.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*MediaStreamAudioSourceNode) SetChannelCountMode(channelCountMode *audio.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*MediaStreamAudioSourceNode) ChannelInterpretation() (channelInterpretation *audio.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*MediaStreamAudioSourceNode) SetChannelInterpretation(channelInterpretation *audio.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*MediaStreamAudioSourceNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*MediaStreamAudioSourceNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*MediaStreamAudioSourceNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}

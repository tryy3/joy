package audionode

import (
	"github.com/matthewmueller/joy/dom/audio"
	"github.com/matthewmueller/joy/dom/window"
)

// AudioNode interface
// js:"AudioNode"
type AudioNode interface {
	window.EventTarget

	// Connect
	// js:"connect"
	// jsrewrite:"$_.connect($1, $2, $3)"
	Connect(destination AudioNode, output *uint, input *uint) (a AudioNode)

	// Disconnect
	// js:"disconnect"
	// jsrewrite:"$_.disconnect()"
	Disconnect()

	// ChannelCount prop
	// js:"channelCount"
	// jsrewrite:"$_.channelCount"
	ChannelCount() (channelCount uint)

	// SetChannelCount prop
	// js:"channelCount"
	// jsrewrite:"$_.channelCount = $1"
	SetChannelCount(channelCount uint)

	// ChannelCountMode prop
	// js:"channelCountMode"
	// jsrewrite:"$_.channelCountMode"
	ChannelCountMode() (channelCountMode *audio.ChannelCountMode)

	// SetChannelCountMode prop
	// js:"channelCountMode"
	// jsrewrite:"$_.channelCountMode = $1"
	SetChannelCountMode(channelCountMode *audio.ChannelCountMode)

	// ChannelInterpretation prop
	// js:"channelInterpretation"
	// jsrewrite:"$_.channelInterpretation"
	ChannelInterpretation() (channelInterpretation *audio.ChannelInterpretation)

	// SetChannelInterpretation prop
	// js:"channelInterpretation"
	// jsrewrite:"$_.channelInterpretation = $1"
	SetChannelInterpretation(channelInterpretation *audio.ChannelInterpretation)

	// Context prop
	// js:"context"
	// jsrewrite:"$_.context"
	Context() (context AudioContext)

	// NumberOfInputs prop
	// js:"numberOfInputs"
	// jsrewrite:"$_.numberOfInputs"
	NumberOfInputs() (numberOfInputs uint)

	// NumberOfOutputs prop
	// js:"numberOfOutputs"
	// jsrewrite:"$_.numberOfOutputs"
	NumberOfOutputs() (numberOfOutputs uint)
}

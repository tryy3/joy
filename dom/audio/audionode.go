package audio

import "github.com/matthewmueller/joy/dom/event"

type AudioNode interface {
	event.
		EventTarget

	Connect(destination AudioNode, output *uint, input *uint) (a AudioNode)

	Disconnect()

	ChannelCount() (channelCount uint)

	SetChannelCount(channelCount uint)

	ChannelCountMode() (channelCountMode *ChannelCountMode)

	SetChannelCountMode(channelCountMode *ChannelCountMode)

	ChannelInterpretation() (channelInterpretation *ChannelInterpretation)

	SetChannelInterpretation(channelInterpretation *ChannelInterpretation)

	Context() (context AudioContext)

	NumberOfInputs() (numberOfInputs uint)

	NumberOfOutputs() (numberOfOutputs uint)
}

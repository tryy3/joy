package audio

import "github.com/matthewmueller/joy/macro"

type AudioBuffer struct {
}

func (*AudioBuffer) CopyFromChannel(destination []float32, channelNumber int, startInChannel *uint) {
	macro.Rewrite("$_.copyFromChannel($1, $2, $3)", destination, channelNumber, startInChannel)
}

func (*AudioBuffer) CopyToChannel(source []float32, channelNumber int, startInChannel *uint) {
	macro.Rewrite("$_.copyToChannel($1, $2, $3)", source, channelNumber, startInChannel)
}

func (*AudioBuffer) GetChannelData(channel uint) (f []float32) {
	macro.Rewrite("$_.getChannelData($1)", channel)
	return f
}

func (*AudioBuffer) Duration() (duration float32) {
	macro.Rewrite("$_.duration")
	return duration
}

func (*AudioBuffer) Length() (length int) {
	macro.Rewrite("$_.length")
	return length
}

func (*AudioBuffer) NumberOfChannels() (numberOfChannels int) {
	macro.Rewrite("$_.numberOfChannels")
	return numberOfChannels
}

func (*AudioBuffer) SampleRate() (sampleRate float32) {
	macro.Rewrite("$_.sampleRate")
	return sampleRate
}

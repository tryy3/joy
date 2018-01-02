package ms

type MSAudioSendPayload struct {
	*MSPayloadBase

	audioFECUsed	*bool
	samplingRate	*uint
	sendMutePercent	*float32
	signal		*MSAudioSendSignal
}

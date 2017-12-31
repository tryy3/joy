package ms


import "github.com/matthewmueller/joy/dom/mspayloadbase"

type MSVideoPayload struct {
	*mspayloadbase.MSPayloadBase

	durationSeconds     *float32
	resolution          *string
	videoBitRateAvg     *uint
	videoBitRateMax     *uint
	videoFrameRateAvg   *float32
	videoPacketLossRate *float32
}

package msaudiosendpayload

import (
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/mspayloadbase"
)

type MSAudioSendPayload struct {
	*mspayloadbase.MSPayloadBase

	audioFECUsed    *bool
	samplingRate    *uint
	sendMutePercent *float32
	signal          *ms.MSAudioSendSignal
}

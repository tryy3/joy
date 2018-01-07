package speechsynthesiseventinit

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
)

type SpeechSynthesisEventInit struct {
	*event.EventInit

	charIndex   *uint
	elapsedTime *float32
	name        *string
	utterance   *window.SpeechSynthesisUtterance
}

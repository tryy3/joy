package speech

import "github.com/matthewmueller/joy/dom/event"

type SpeechSynthesisEventInit struct {
	*event.EventInit

	charIndex	*uint
	elapsedTime	*float32
	name		*string
	utterance	*SpeechSynthesisUtterance
}

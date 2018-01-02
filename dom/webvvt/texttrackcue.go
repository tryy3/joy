package webvvt

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/document"
)

type TextTrackCue interface {
	event.
		EventTarget

	GetCueAsHTML() (w *document.DocumentFragment)

	EndTime() (endTime float32)

	SetEndTime(endTime float32)

	ID() (id string)

	SetID(id string)

	Onenter() (onenter func(event.Event))

	SetOnenter(onenter func(event.Event))

	Onexit() (onexit func(event.Event))

	SetOnexit(onexit func(event.Event))

	PauseOnExit() (pauseOnExit bool)

	SetPauseOnExit(pauseOnExit bool)

	StartTime() (startTime float32)

	SetStartTime(startTime float32)

	Text() (text string)

	SetText(text string)

	Track() (track *TextTrack)
}

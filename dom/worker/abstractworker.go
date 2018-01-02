package worker

import "github.com/matthewmueller/joy/dom/event"

type AbstractWorker interface {
	Onerror() (onerror func(event.Event))

	SetOnerror(onerror func(event.Event))
}

package document

import "github.com/matthewmueller/joy/dom/event"

type DocumentEvent interface {
	CreateEvent(eventInterface string) (e event.Event)
}

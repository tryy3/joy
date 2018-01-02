package navigation

import "github.com/matthewmueller/joy/dom/event"

type NavigationEvent interface {
	event.
		Event

	URI() (uri string)
}

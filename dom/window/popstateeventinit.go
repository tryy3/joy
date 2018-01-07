package window

import "github.com/matthewmueller/joy/dom/event"

type PopStateEventInit struct {
	*event.EventInit

	state *interface{}
}

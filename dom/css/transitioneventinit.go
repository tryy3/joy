package css

import "github.com/matthewmueller/joy/dom/event"

type TransitionEventInit struct {
	*event.EventInit

	elapsedTime	*float32
	propertyName	*string
}

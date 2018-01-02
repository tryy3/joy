package dom

import "github.com/matthewmueller/joy/dom/event"

type FocusNavigationEventInit struct {
	*event.EventInit

	navigationReason	*string
	originHeight		*float32
	originLeft		*float32
	originTop		*float32
	originWidth		*float32
}

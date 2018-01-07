package navigationeventwithreferrer

import "github.com/matthewmueller/joy/dom/navigationevent"

// NavigationEventWithReferrer interface
// js:"NavigationEventWithReferrer"
type NavigationEventWithReferrer interface {
	navigationevent.NavigationEvent

	// Referer prop
	// js:"referer"
	// jsrewrite:"$_.referer"
	Referer() (referer string)
}

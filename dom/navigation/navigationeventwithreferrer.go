package navigation


// NavigationEventWithReferrer interface
// js:"NavigationEventWithReferrer"
type NavigationEventWithReferrer interface {
	NavigationEvent

	// Referer prop
	// js:"referer"
	// jsrewrite:"$_.referer"
	Referer() (referer string)
}

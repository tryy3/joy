package navigation

type NavigationEventWithReferrer interface {
	NavigationEvent

	Referer() (referer string)
}

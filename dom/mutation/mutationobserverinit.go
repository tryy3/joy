package mutation

type MutationObserverInit struct {
	attributeFilter       *[]string
	attributeOldValue     *bool
	attributes            *bool
	characterData         *bool
	characterDataOldValue *bool
	childList             *bool
	subtree               *bool
}

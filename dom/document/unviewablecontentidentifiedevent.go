package document

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/navigation"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ navigation.NavigationEventWithReferrer = (*UnviewableContentIdentifiedEvent)(nil)
var _ navigation.NavigationEvent = (*UnviewableContentIdentifiedEvent)(nil)
var _ event.Event = (*UnviewableContentIdentifiedEvent)(nil)

type UnviewableContentIdentifiedEvent struct {
}

func (*UnviewableContentIdentifiedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*UnviewableContentIdentifiedEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*UnviewableContentIdentifiedEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*UnviewableContentIdentifiedEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*UnviewableContentIdentifiedEvent) MediaType() (mediaType string) {
	macro.Rewrite("$_.mediaType")
	return mediaType
}

func (*UnviewableContentIdentifiedEvent) Referer() (referer string) {
	macro.Rewrite("$_.referer")
	return referer
}

func (*UnviewableContentIdentifiedEvent) URI() (uri string) {
	macro.Rewrite("$_.uri")
	return uri
}

func (*UnviewableContentIdentifiedEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*UnviewableContentIdentifiedEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*UnviewableContentIdentifiedEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*UnviewableContentIdentifiedEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*UnviewableContentIdentifiedEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*UnviewableContentIdentifiedEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*UnviewableContentIdentifiedEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*UnviewableContentIdentifiedEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*UnviewableContentIdentifiedEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*UnviewableContentIdentifiedEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*UnviewableContentIdentifiedEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*UnviewableContentIdentifiedEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*UnviewableContentIdentifiedEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*UnviewableContentIdentifiedEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

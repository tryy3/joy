package navigation

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ NavigationEvent = (*NavigationCompletedEvent)(nil)
var _ event.Event = (*NavigationCompletedEvent)(nil)

type NavigationCompletedEvent struct {
}

func (*NavigationCompletedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*NavigationCompletedEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*NavigationCompletedEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*NavigationCompletedEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*NavigationCompletedEvent) IsSuccess() (isSuccess bool) {
	macro.Rewrite("$_.isSuccess")
	return isSuccess
}

func (*NavigationCompletedEvent) WebErrorStatus() (webErrorStatus uint64) {
	macro.Rewrite("$_.webErrorStatus")
	return webErrorStatus
}

func (*NavigationCompletedEvent) URI() (uri string) {
	macro.Rewrite("$_.uri")
	return uri
}

func (*NavigationCompletedEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*NavigationCompletedEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*NavigationCompletedEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*NavigationCompletedEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*NavigationCompletedEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*NavigationCompletedEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*NavigationCompletedEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*NavigationCompletedEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*NavigationCompletedEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*NavigationCompletedEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*NavigationCompletedEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*NavigationCompletedEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*NavigationCompletedEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*NavigationCompletedEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

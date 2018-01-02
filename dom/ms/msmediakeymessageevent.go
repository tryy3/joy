package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*MSMediaKeyMessageEvent)(nil)

type MSMediaKeyMessageEvent struct {
}

func (*MSMediaKeyMessageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*MSMediaKeyMessageEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*MSMediaKeyMessageEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*MSMediaKeyMessageEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*MSMediaKeyMessageEvent) DestinationURL() (destinationURL string) {
	macro.Rewrite("$_.destinationURL")
	return destinationURL
}

func (*MSMediaKeyMessageEvent) Message() (message []uint8) {
	macro.Rewrite("$_.message")
	return message
}

func (*MSMediaKeyMessageEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*MSMediaKeyMessageEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*MSMediaKeyMessageEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*MSMediaKeyMessageEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*MSMediaKeyMessageEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*MSMediaKeyMessageEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*MSMediaKeyMessageEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*MSMediaKeyMessageEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*MSMediaKeyMessageEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*MSMediaKeyMessageEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*MSMediaKeyMessageEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*MSMediaKeyMessageEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*MSMediaKeyMessageEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*MSMediaKeyMessageEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

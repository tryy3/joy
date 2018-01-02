package event

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
)

var _ Event = (*CustomEvent)(nil)

func New(typearg string, eventinitdict *CustomEventInit) *CustomEvent {
	macro.Rewrite("new CustomEvent($1, $2)", typearg, eventinitdict)
	return &CustomEvent{}
}

type CustomEvent struct {
}

func (*CustomEvent) InitCustomEvent(typeArg string, canBubbleArg bool, cancelableArg bool, detailArg interface{}) {
	macro.Rewrite("$_.initCustomEvent($1, $2, $3, $4)", typeArg, canBubbleArg, cancelableArg, detailArg)
}

func (*CustomEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*CustomEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*CustomEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*CustomEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*CustomEvent) Detail() (detail interface{}) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*CustomEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*CustomEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*CustomEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*CustomEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*CustomEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*CustomEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*CustomEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*CustomEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*CustomEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*CustomEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*CustomEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*CustomEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*CustomEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*CustomEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

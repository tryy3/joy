package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*PopStateEvent)(nil)

func NewPopStateEvent(typearg string, eventinitdict *window.PopStateEventInit) *PopStateEvent {
	macro.Rewrite("new PopStateEvent($1, $2)", typearg, eventinitdict)
	return &PopStateEvent{}
}

type PopStateEvent struct {
}

func (*PopStateEvent) InitPopStateEvent(typeArg string, canBubbleArg bool, cancelableArg bool, stateArg interface{}) {
	macro.Rewrite("$_.initPopStateEvent($1, $2, $3, $4)", typeArg, canBubbleArg, cancelableArg, stateArg)
}

func (*PopStateEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*PopStateEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*PopStateEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*PopStateEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*PopStateEvent) State() (state interface{}) {
	macro.Rewrite("$_.state")
	return state
}

func (*PopStateEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*PopStateEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*PopStateEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*PopStateEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*PopStateEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*PopStateEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*PopStateEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*PopStateEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*PopStateEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*PopStateEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*PopStateEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*PopStateEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*PopStateEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*PopStateEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

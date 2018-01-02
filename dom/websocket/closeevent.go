package websocket

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*CloseEvent)(nil)

func New(typearg string, eventinitdict *CloseEventInit) *CloseEvent {
	macro.Rewrite("new CloseEvent($1, $2)", typearg, eventinitdict)
	return &CloseEvent{}
}

type CloseEvent struct {
}

func (*CloseEvent) InitCloseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, wasCleanArg bool, codeArg uint8, reasonArg string) {
	macro.Rewrite("$_.initCloseEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, wasCleanArg, codeArg, reasonArg)
}

func (*CloseEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*CloseEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*CloseEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*CloseEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*CloseEvent) Code() (code uint8) {
	macro.Rewrite("$_.code")
	return code
}

func (*CloseEvent) Reason() (reason string) {
	macro.Rewrite("$_.reason")
	return reason
}

func (*CloseEvent) WasClean() (wasClean bool) {
	macro.Rewrite("$_.wasClean")
	return wasClean
}

func (*CloseEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*CloseEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*CloseEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*CloseEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*CloseEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*CloseEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*CloseEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*CloseEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*CloseEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*CloseEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*CloseEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*CloseEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*CloseEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*CloseEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

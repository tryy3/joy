package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*LongRunningScriptDetectedEvent)(nil)

type LongRunningScriptDetectedEvent struct {
}

func (*LongRunningScriptDetectedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*LongRunningScriptDetectedEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*LongRunningScriptDetectedEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*LongRunningScriptDetectedEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*LongRunningScriptDetectedEvent) ExecutionTime() (executionTime int) {
	macro.Rewrite("$_.executionTime")
	return executionTime
}

func (*LongRunningScriptDetectedEvent) StopPageScriptExecution() (stopPageScriptExecution bool) {
	macro.Rewrite("$_.stopPageScriptExecution")
	return stopPageScriptExecution
}

func (*LongRunningScriptDetectedEvent) SetStopPageScriptExecution(stopPageScriptExecution bool) {
	macro.Rewrite("$_.stopPageScriptExecution = $1", stopPageScriptExecution)
}

func (*LongRunningScriptDetectedEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*LongRunningScriptDetectedEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*LongRunningScriptDetectedEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*LongRunningScriptDetectedEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*LongRunningScriptDetectedEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*LongRunningScriptDetectedEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*LongRunningScriptDetectedEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*LongRunningScriptDetectedEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*LongRunningScriptDetectedEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*LongRunningScriptDetectedEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*LongRunningScriptDetectedEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*LongRunningScriptDetectedEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*LongRunningScriptDetectedEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*LongRunningScriptDetectedEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

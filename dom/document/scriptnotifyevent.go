package document

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*ScriptNotifyEvent)(nil)

type ScriptNotifyEvent struct {
}

func (*ScriptNotifyEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*ScriptNotifyEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*ScriptNotifyEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*ScriptNotifyEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*ScriptNotifyEvent) CallingURI() (callingUri string) {
	macro.Rewrite("$_.callingUri")
	return callingUri
}

func (*ScriptNotifyEvent) Value() (value string) {
	macro.Rewrite("$_.value")
	return value
}

func (*ScriptNotifyEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*ScriptNotifyEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*ScriptNotifyEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*ScriptNotifyEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*ScriptNotifyEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*ScriptNotifyEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*ScriptNotifyEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*ScriptNotifyEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*ScriptNotifyEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*ScriptNotifyEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*ScriptNotifyEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*ScriptNotifyEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*ScriptNotifyEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*ScriptNotifyEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

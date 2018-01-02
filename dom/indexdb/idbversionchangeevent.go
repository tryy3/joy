package indexdb

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*IDBVersionChangeEvent)(nil)

type IDBVersionChangeEvent struct {
}

func (*IDBVersionChangeEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*IDBVersionChangeEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*IDBVersionChangeEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*IDBVersionChangeEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*IDBVersionChangeEvent) NewVersion() (newVersion uint64) {
	macro.Rewrite("$_.newVersion")
	return newVersion
}

func (*IDBVersionChangeEvent) OldVersion() (oldVersion uint64) {
	macro.Rewrite("$_.oldVersion")
	return oldVersion
}

func (*IDBVersionChangeEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*IDBVersionChangeEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*IDBVersionChangeEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*IDBVersionChangeEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*IDBVersionChangeEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*IDBVersionChangeEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*IDBVersionChangeEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*IDBVersionChangeEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*IDBVersionChangeEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*IDBVersionChangeEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*IDBVersionChangeEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*IDBVersionChangeEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*IDBVersionChangeEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*IDBVersionChangeEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

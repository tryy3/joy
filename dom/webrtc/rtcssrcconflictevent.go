package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*RTCSsrcConflictEvent)(nil)

type RTCSsrcConflictEvent struct {
}

func (*RTCSsrcConflictEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*RTCSsrcConflictEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*RTCSsrcConflictEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*RTCSsrcConflictEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*RTCSsrcConflictEvent) Ssrc() (ssrc uint) {
	macro.Rewrite("$_.ssrc")
	return ssrc
}

func (*RTCSsrcConflictEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*RTCSsrcConflictEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*RTCSsrcConflictEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*RTCSsrcConflictEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*RTCSsrcConflictEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*RTCSsrcConflictEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*RTCSsrcConflictEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*RTCSsrcConflictEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*RTCSsrcConflictEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*RTCSsrcConflictEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*RTCSsrcConflictEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*RTCSsrcConflictEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*RTCSsrcConflictEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*RTCSsrcConflictEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

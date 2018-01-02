package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*RTCIceGathererEvent)(nil)

type RTCIceGathererEvent struct {
}

func (*RTCIceGathererEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*RTCIceGathererEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*RTCIceGathererEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*RTCIceGathererEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*RTCIceGathererEvent) Candidate() (candidate interface{}) {
	macro.Rewrite("$_.candidate")
	return candidate
}

func (*RTCIceGathererEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*RTCIceGathererEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*RTCIceGathererEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*RTCIceGathererEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*RTCIceGathererEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*RTCIceGathererEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*RTCIceGathererEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*RTCIceGathererEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*RTCIceGathererEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*RTCIceGathererEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*RTCIceGathererEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*RTCIceGathererEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*RTCIceGathererEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*RTCIceGathererEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*RTCDtlsTransportStateChangedEvent)(nil)

type RTCDtlsTransportStateChangedEvent struct {
}

func (*RTCDtlsTransportStateChangedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*RTCDtlsTransportStateChangedEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*RTCDtlsTransportStateChangedEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*RTCDtlsTransportStateChangedEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*RTCDtlsTransportStateChangedEvent) State() (state *RTCDtlsTransportState) {
	macro.Rewrite("$_.state")
	return state
}

func (*RTCDtlsTransportStateChangedEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*RTCDtlsTransportStateChangedEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*RTCDtlsTransportStateChangedEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*RTCDtlsTransportStateChangedEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*RTCDtlsTransportStateChangedEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*RTCDtlsTransportStateChangedEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*RTCDtlsTransportStateChangedEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*RTCDtlsTransportStateChangedEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*RTCDtlsTransportStateChangedEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*RTCDtlsTransportStateChangedEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*RTCDtlsTransportStateChangedEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*RTCDtlsTransportStateChangedEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*RTCDtlsTransportStateChangedEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*RTCDtlsTransportStateChangedEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

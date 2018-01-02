package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*RTCIceCandidatePairChangedEvent)(nil)

type RTCIceCandidatePairChangedEvent struct {
}

func (*RTCIceCandidatePairChangedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*RTCIceCandidatePairChangedEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*RTCIceCandidatePairChangedEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*RTCIceCandidatePairChangedEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*RTCIceCandidatePairChangedEvent) Pair() (pair *RTCIceCandidatePair) {
	macro.Rewrite("$_.pair")
	return pair
}

func (*RTCIceCandidatePairChangedEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*RTCIceCandidatePairChangedEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*RTCIceCandidatePairChangedEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*RTCIceCandidatePairChangedEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*RTCIceCandidatePairChangedEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*RTCIceCandidatePairChangedEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*RTCIceCandidatePairChangedEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*RTCIceCandidatePairChangedEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*RTCIceCandidatePairChangedEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*RTCIceCandidatePairChangedEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*RTCIceCandidatePairChangedEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*RTCIceCandidatePairChangedEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*RTCIceCandidatePairChangedEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*RTCIceCandidatePairChangedEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

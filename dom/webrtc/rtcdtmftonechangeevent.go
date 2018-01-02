package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*RTCDTMFToneChangeEvent)(nil)

func New(typearg string, eventinitdict *RTCDTMFToneChangeEventInit) *RTCDTMFToneChangeEvent {
	macro.Rewrite("new RTCDTMFToneChangeEvent($1, $2)", typearg, eventinitdict)
	return &RTCDTMFToneChangeEvent{}
}

type RTCDTMFToneChangeEvent struct {
}

func (*RTCDTMFToneChangeEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*RTCDTMFToneChangeEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*RTCDTMFToneChangeEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*RTCDTMFToneChangeEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*RTCDTMFToneChangeEvent) Tone() (tone string) {
	macro.Rewrite("$_.tone")
	return tone
}

func (*RTCDTMFToneChangeEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*RTCDTMFToneChangeEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*RTCDTMFToneChangeEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*RTCDTMFToneChangeEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*RTCDTMFToneChangeEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*RTCDTMFToneChangeEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*RTCDTMFToneChangeEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*RTCDTMFToneChangeEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*RTCDTMFToneChangeEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*RTCDTMFToneChangeEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*RTCDTMFToneChangeEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*RTCDTMFToneChangeEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*RTCDTMFToneChangeEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*RTCDTMFToneChangeEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

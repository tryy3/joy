package webrtc

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*RTCPeerConnectionIceEvent)(nil)

func New(kind string, eventinitdict *RTCPeerConnectionIceEventInit) *RTCPeerConnectionIceEvent {
	macro.Rewrite("new RTCPeerConnectionIceEvent($1, $2)", kind, eventinitdict)
	return &RTCPeerConnectionIceEvent{}
}

type RTCPeerConnectionIceEvent struct {
}

func (*RTCPeerConnectionIceEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*RTCPeerConnectionIceEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*RTCPeerConnectionIceEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*RTCPeerConnectionIceEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*RTCPeerConnectionIceEvent) Candidate() (candidate *RTCIceCandidate) {
	macro.Rewrite("$_.candidate")
	return candidate
}

func (*RTCPeerConnectionIceEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*RTCPeerConnectionIceEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*RTCPeerConnectionIceEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*RTCPeerConnectionIceEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*RTCPeerConnectionIceEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*RTCPeerConnectionIceEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*RTCPeerConnectionIceEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*RTCPeerConnectionIceEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*RTCPeerConnectionIceEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*RTCPeerConnectionIceEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*RTCPeerConnectionIceEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*RTCPeerConnectionIceEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*RTCPeerConnectionIceEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*RTCPeerConnectionIceEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

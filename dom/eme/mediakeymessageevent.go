package eme

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*MediaKeyMessageEvent)(nil)

func New(kind string, eventinitdict *MediaKeyMessageEventInit) *MediaKeyMessageEvent {
	macro.Rewrite("new MediaKeyMessageEvent($1, $2)", kind, eventinitdict)
	return &MediaKeyMessageEvent{}
}

type MediaKeyMessageEvent struct {
}

func (*MediaKeyMessageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*MediaKeyMessageEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*MediaKeyMessageEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*MediaKeyMessageEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*MediaKeyMessageEvent) Message() (message []byte) {
	macro.Rewrite("$_.message")
	return message
}

func (*MediaKeyMessageEvent) MessageType() (messageType *MediaKeyMessageType) {
	macro.Rewrite("$_.messageType")
	return messageType
}

func (*MediaKeyMessageEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*MediaKeyMessageEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*MediaKeyMessageEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*MediaKeyMessageEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*MediaKeyMessageEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*MediaKeyMessageEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*MediaKeyMessageEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*MediaKeyMessageEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*MediaKeyMessageEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*MediaKeyMessageEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*MediaKeyMessageEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*MediaKeyMessageEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*MediaKeyMessageEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*MediaKeyMessageEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

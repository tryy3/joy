package event

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
)

var _ Event = (*HashChangeEvent)(nil)

func NewHashChangeEvent(typearg string, eventinitdict *HashChangeEventInit) *HashChangeEvent {
	macro.Rewrite("new HashChangeEvent($1, $2)", typearg, eventinitdict)
	return &HashChangeEvent{}
}

type HashChangeEvent struct {
}

func (*HashChangeEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*HashChangeEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*HashChangeEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*HashChangeEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*HashChangeEvent) NewURL() (newURL string) {
	macro.Rewrite("$_.newURL")
	return newURL
}

func (*HashChangeEvent) OldURL() (oldURL string) {
	macro.Rewrite("$_.oldURL")
	return oldURL
}

func (*HashChangeEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*HashChangeEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*HashChangeEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*HashChangeEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*HashChangeEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*HashChangeEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*HashChangeEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*HashChangeEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*HashChangeEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*HashChangeEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*HashChangeEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*HashChangeEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*HashChangeEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*HashChangeEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

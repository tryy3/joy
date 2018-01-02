package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/device"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*DeviceLightEvent)(nil)

func NewDeviceLightEvent(typearg string, eventinitdict *device.DeviceLightEventInit) *DeviceLightEvent {
	macro.Rewrite("new DeviceLightEvent($1, $2)", typearg, eventinitdict)
	return &DeviceLightEvent{}
}

type DeviceLightEvent struct {
}

func (*DeviceLightEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*DeviceLightEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*DeviceLightEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*DeviceLightEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*DeviceLightEvent) Value() (value float32) {
	macro.Rewrite("$_.value")
	return value
}

func (*DeviceLightEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*DeviceLightEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*DeviceLightEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*DeviceLightEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*DeviceLightEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*DeviceLightEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*DeviceLightEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*DeviceLightEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*DeviceLightEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*DeviceLightEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*DeviceLightEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*DeviceLightEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*DeviceLightEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*DeviceLightEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

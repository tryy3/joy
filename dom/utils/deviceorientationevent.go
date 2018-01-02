package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/device"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*DeviceOrientationEvent)(nil)

func NewDeviceOrientationEvent(typearg string, eventinitdict *device.DeviceOrientationEventInit) *DeviceOrientationEvent {
	macro.Rewrite("new DeviceOrientationEvent($1, $2)", typearg, eventinitdict)
	return &DeviceOrientationEvent{}
}

type DeviceOrientationEvent struct {
}

func (*DeviceOrientationEvent) InitDeviceOrientationEvent(kind string, bubbles bool, cancelable bool, alpha float32, beta float32, gamma float32, absolute bool) {
	macro.Rewrite("$_.initDeviceOrientationEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, alpha, beta, gamma, absolute)
}

func (*DeviceOrientationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*DeviceOrientationEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*DeviceOrientationEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*DeviceOrientationEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*DeviceOrientationEvent) Absolute() (absolute bool) {
	macro.Rewrite("$_.absolute")
	return absolute
}

func (*DeviceOrientationEvent) Alpha() (alpha float32) {
	macro.Rewrite("$_.alpha")
	return alpha
}

func (*DeviceOrientationEvent) Beta() (beta float32) {
	macro.Rewrite("$_.beta")
	return beta
}

func (*DeviceOrientationEvent) Gamma() (gamma float32) {
	macro.Rewrite("$_.gamma")
	return gamma
}

func (*DeviceOrientationEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*DeviceOrientationEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*DeviceOrientationEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*DeviceOrientationEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*DeviceOrientationEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*DeviceOrientationEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*DeviceOrientationEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*DeviceOrientationEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*DeviceOrientationEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*DeviceOrientationEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*DeviceOrientationEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*DeviceOrientationEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*DeviceOrientationEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*DeviceOrientationEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

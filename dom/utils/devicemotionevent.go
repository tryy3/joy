package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/device"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*DeviceMotionEvent)(nil)

func NewDeviceMotionEvent(typearg string, eventinitdict *device.DeviceMotionEventInit) *DeviceMotionEvent {
	macro.Rewrite("new DeviceMotionEvent($1, $2)", typearg, eventinitdict)
	return &DeviceMotionEvent{}
}

type DeviceMotionEvent struct {
}

func (*DeviceMotionEvent) InitDeviceMotionEvent(kind string, bubbles bool, cancelable bool, acceleration *device.DeviceAccelerationDict, accelerationIncludingGravity *device.DeviceAccelerationDict, rotationRate *device.DeviceRotationRateDict, interval float32) {
	macro.Rewrite("$_.initDeviceMotionEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, acceleration, accelerationIncludingGravity, rotationRate, interval)
}

func (*DeviceMotionEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*DeviceMotionEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*DeviceMotionEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*DeviceMotionEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*DeviceMotionEvent) Acceleration() (acceleration *device.DeviceAcceleration) {
	macro.Rewrite("$_.acceleration")
	return acceleration
}

func (*DeviceMotionEvent) AccelerationIncludingGravity() (accelerationIncludingGravity *device.DeviceAcceleration) {
	macro.Rewrite("$_.accelerationIncludingGravity")
	return accelerationIncludingGravity
}

func (*DeviceMotionEvent) Interval() (interval float32) {
	macro.Rewrite("$_.interval")
	return interval
}

func (*DeviceMotionEvent) RotationRate() (rotationRate *device.DeviceRotationRate) {
	macro.Rewrite("$_.rotationRate")
	return rotationRate
}

func (*DeviceMotionEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*DeviceMotionEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*DeviceMotionEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*DeviceMotionEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*DeviceMotionEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*DeviceMotionEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*DeviceMotionEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*DeviceMotionEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*DeviceMotionEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*DeviceMotionEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*DeviceMotionEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*DeviceMotionEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*DeviceMotionEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*DeviceMotionEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

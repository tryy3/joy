package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ ui.UIEvent = (*MSManipulationEvent)(nil)
var _ event.Event = (*MSManipulationEvent)(nil)

type MSManipulationEvent struct {
}

func (*MSManipulationEvent) InitMSManipulationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, lastState int, currentState int) {
	macro.Rewrite("$_.initMSManipulationEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, lastState, currentState)
}

func (*MSManipulationEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*MSManipulationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*MSManipulationEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*MSManipulationEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*MSManipulationEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*MSManipulationEvent) CurrentState() (currentState int) {
	macro.Rewrite("$_.currentState")
	return currentState
}

func (*MSManipulationEvent) InertiaDestinationX() (inertiaDestinationX int) {
	macro.Rewrite("$_.inertiaDestinationX")
	return inertiaDestinationX
}

func (*MSManipulationEvent) InertiaDestinationY() (inertiaDestinationY int) {
	macro.Rewrite("$_.inertiaDestinationY")
	return inertiaDestinationY
}

func (*MSManipulationEvent) LastState() (lastState int) {
	macro.Rewrite("$_.lastState")
	return lastState
}

func (*MSManipulationEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*MSManipulationEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*MSManipulationEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*MSManipulationEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*MSManipulationEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*MSManipulationEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*MSManipulationEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*MSManipulationEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*MSManipulationEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*MSManipulationEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*MSManipulationEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*MSManipulationEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*MSManipulationEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*MSManipulationEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*MSManipulationEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*MSManipulationEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

package ms


import "github.com/matthewmueller/joy/macro"

var _ UIEvent = (*MSManipulationEvent)(nil)
var _ Event = (*MSManipulationEvent)(nil)

// MSManipulationEvent struct
// js:"MSManipulationEvent,omit"
type MSManipulationEvent struct {
}

// InitMSManipulationEvent fn
// js:"initMSManipulationEvent"
func (*MSManipulationEvent) InitMSManipulationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, lastState int, currentState int) {
	macro.Rewrite("$_.initMSManipulationEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, lastState, currentState)
}

// InitUIEvent fn
// js:"initUIEvent"
func (*MSManipulationEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*MSManipulationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MSManipulationEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MSManipulationEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MSManipulationEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// CurrentState prop
// js:"currentState"
func (*MSManipulationEvent) CurrentState() (currentState int) {
	macro.Rewrite("$_.currentState")
	return currentState
}

// InertiaDestinationX prop
// js:"inertiaDestinationX"
func (*MSManipulationEvent) InertiaDestinationX() (inertiaDestinationX int) {
	macro.Rewrite("$_.inertiaDestinationX")
	return inertiaDestinationX
}

// InertiaDestinationY prop
// js:"inertiaDestinationY"
func (*MSManipulationEvent) InertiaDestinationY() (inertiaDestinationY int) {
	macro.Rewrite("$_.inertiaDestinationY")
	return inertiaDestinationY
}

// LastState prop
// js:"lastState"
func (*MSManipulationEvent) LastState() (lastState int) {
	macro.Rewrite("$_.lastState")
	return lastState
}

// Detail prop
// js:"detail"
func (*MSManipulationEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*MSManipulationEvent) View() (view *Window) {
	macro.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*MSManipulationEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MSManipulationEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MSManipulationEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MSManipulationEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MSManipulationEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MSManipulationEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MSManipulationEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MSManipulationEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MSManipulationEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MSManipulationEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MSManipulationEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MSManipulationEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MSManipulationEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MSManipulationEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

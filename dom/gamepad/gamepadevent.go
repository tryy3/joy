package gamepad

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*GamepadEvent)(nil)

func New(typearg string, eventinitdict *GamepadEventInit) *GamepadEvent {
	macro.Rewrite("new GamepadEvent($1, $2)", typearg, eventinitdict)
	return &GamepadEvent{}
}

type GamepadEvent struct {
}

func (*GamepadEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*GamepadEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*GamepadEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*GamepadEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*GamepadEvent) Gamepad() (gamepad *Gamepad) {
	macro.Rewrite("$_.gamepad")
	return gamepad
}

func (*GamepadEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*GamepadEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*GamepadEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*GamepadEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*GamepadEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*GamepadEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*GamepadEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*GamepadEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*GamepadEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*GamepadEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*GamepadEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*GamepadEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*GamepadEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*GamepadEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

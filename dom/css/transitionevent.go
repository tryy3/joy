package css

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*TransitionEvent)(nil)

func New(typearg string, eventinitdict *TransitionEventInit) *TransitionEvent {
	macro.Rewrite("new TransitionEvent($1, $2)", typearg, eventinitdict)
	return &TransitionEvent{}
}

type TransitionEvent struct {
}

func (*TransitionEvent) InitTransitionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, propertyNameArg string, elapsedTimeArg float32) {
	macro.Rewrite("$_.initTransitionEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, propertyNameArg, elapsedTimeArg)
}

func (*TransitionEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*TransitionEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*TransitionEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*TransitionEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*TransitionEvent) ElapsedTime() (elapsedTime float32) {
	macro.Rewrite("$_.elapsedTime")
	return elapsedTime
}

func (*TransitionEvent) PropertyName() (propertyName string) {
	macro.Rewrite("$_.propertyName")
	return propertyName
}

func (*TransitionEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*TransitionEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*TransitionEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*TransitionEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*TransitionEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*TransitionEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*TransitionEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*TransitionEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*TransitionEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*TransitionEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*TransitionEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*TransitionEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*TransitionEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*TransitionEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

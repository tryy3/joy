package animation

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*AnimationEvent)(nil)

func New(typearg string, eventinitdict *AnimationEventInit) *AnimationEvent {
	macro.Rewrite("new AnimationEvent($1, $2)", typearg, eventinitdict)
	return &AnimationEvent{}
}

type AnimationEvent struct {
}

func (*AnimationEvent) InitAnimationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, animationNameArg string, elapsedTimeArg float32) {
	macro.Rewrite("$_.initAnimationEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, animationNameArg, elapsedTimeArg)
}

func (*AnimationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*AnimationEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*AnimationEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*AnimationEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*AnimationEvent) AnimationName() (animationName string) {
	macro.Rewrite("$_.animationName")
	return animationName
}

func (*AnimationEvent) ElapsedTime() (elapsedTime float32) {
	macro.Rewrite("$_.elapsedTime")
	return elapsedTime
}

func (*AnimationEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*AnimationEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*AnimationEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*AnimationEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*AnimationEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*AnimationEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*AnimationEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*AnimationEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*AnimationEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*AnimationEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*AnimationEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*AnimationEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*AnimationEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*AnimationEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

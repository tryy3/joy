package webgl

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*WebGLContextEvent)(nil)

func New(typearg string, eventinitdict *WebGLContextEventInit) *WebGLContextEvent {
	macro.Rewrite("new WebGLContextEvent($1, $2)", typearg, eventinitdict)
	return &WebGLContextEvent{}
}

type WebGLContextEvent struct {
}

func (*WebGLContextEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*WebGLContextEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*WebGLContextEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*WebGLContextEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*WebGLContextEvent) StatusMessage() (statusMessage string) {
	macro.Rewrite("$_.statusMessage")
	return statusMessage
}

func (*WebGLContextEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*WebGLContextEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*WebGLContextEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*WebGLContextEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*WebGLContextEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*WebGLContextEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*WebGLContextEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*WebGLContextEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*WebGLContextEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*WebGLContextEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*WebGLContextEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*WebGLContextEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*WebGLContextEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*WebGLContextEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

package event

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
)

var _ Event = (*ErrorEvent)(nil)

func New(typearg string, eventinitdict *ErrorEventInit) *ErrorEvent {
	macro.Rewrite("new ErrorEvent($1, $2)", typearg, eventinitdict)
	return &ErrorEvent{}
}

type ErrorEvent struct {
}

func (*ErrorEvent) InitErrorEvent(typeArg string, canBubbleArg bool, cancelableArg bool, messageArg string, filenameArg string, linenoArg uint) {
	macro.Rewrite("$_.initErrorEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, messageArg, filenameArg, linenoArg)
}

func (*ErrorEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*ErrorEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*ErrorEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*ErrorEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*ErrorEvent) Colno() (colno uint) {
	macro.Rewrite("$_.colno")
	return colno
}

func (*ErrorEvent) Error() (err interface{}) {
	macro.Rewrite("$_.error")
	return err
}

func (*ErrorEvent) Filename() (filename string) {
	macro.Rewrite("$_.filename")
	return filename
}

func (*ErrorEvent) Lineno() (lineno uint) {
	macro.Rewrite("$_.lineno")
	return lineno
}

func (*ErrorEvent) Message() (message string) {
	macro.Rewrite("$_.message")
	return message
}

func (*ErrorEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*ErrorEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*ErrorEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*ErrorEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*ErrorEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*ErrorEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*ErrorEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*ErrorEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*ErrorEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*ErrorEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*ErrorEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*ErrorEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*ErrorEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*ErrorEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

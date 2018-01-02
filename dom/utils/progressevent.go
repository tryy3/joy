package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*ProgressEvent)(nil)

func NewProgressEvent(typearg string, eventinitdict *ProgressEventInit) *ProgressEvent {
	macro.Rewrite("new ProgressEvent($1, $2)", typearg, eventinitdict)
	return &ProgressEvent{}
}

type ProgressEvent struct {
}

func (*ProgressEvent) InitProgressEvent(typeArg string, canBubbleArg bool, cancelableArg bool, lengthComputableArg bool, loadedArg uint64, totalArg uint64) {
	macro.Rewrite("$_.initProgressEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, lengthComputableArg, loadedArg, totalArg)
}

func (*ProgressEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*ProgressEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*ProgressEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*ProgressEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*ProgressEvent) LengthComputable() (lengthComputable bool) {
	macro.Rewrite("$_.lengthComputable")
	return lengthComputable
}

func (*ProgressEvent) Loaded() (loaded uint64) {
	macro.Rewrite("$_.loaded")
	return loaded
}

func (*ProgressEvent) Total() (total uint64) {
	macro.Rewrite("$_.total")
	return total
}

func (*ProgressEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*ProgressEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*ProgressEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*ProgressEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*ProgressEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*ProgressEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*ProgressEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*ProgressEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*ProgressEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*ProgressEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*ProgressEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*ProgressEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*ProgressEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*ProgressEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

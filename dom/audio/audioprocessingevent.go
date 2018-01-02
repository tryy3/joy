package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*AudioProcessingEvent)(nil)

type AudioProcessingEvent struct {
}

func (*AudioProcessingEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*AudioProcessingEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*AudioProcessingEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*AudioProcessingEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*AudioProcessingEvent) InputBuffer() (inputBuffer *AudioBuffer) {
	macro.Rewrite("$_.inputBuffer")
	return inputBuffer
}

func (*AudioProcessingEvent) OutputBuffer() (outputBuffer *AudioBuffer) {
	macro.Rewrite("$_.outputBuffer")
	return outputBuffer
}

func (*AudioProcessingEvent) PlaybackTime() (playbackTime float32) {
	macro.Rewrite("$_.playbackTime")
	return playbackTime
}

func (*AudioProcessingEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*AudioProcessingEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*AudioProcessingEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*AudioProcessingEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*AudioProcessingEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*AudioProcessingEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*AudioProcessingEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*AudioProcessingEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*AudioProcessingEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*AudioProcessingEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*AudioProcessingEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*AudioProcessingEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*AudioProcessingEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*AudioProcessingEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

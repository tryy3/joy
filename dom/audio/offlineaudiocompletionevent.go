package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*OfflineAudioCompletionEvent)(nil)

type OfflineAudioCompletionEvent struct {
}

func (*OfflineAudioCompletionEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*OfflineAudioCompletionEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*OfflineAudioCompletionEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*OfflineAudioCompletionEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*OfflineAudioCompletionEvent) RenderedBuffer() (renderedBuffer *AudioBuffer) {
	macro.Rewrite("$_.renderedBuffer")
	return renderedBuffer
}

func (*OfflineAudioCompletionEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*OfflineAudioCompletionEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*OfflineAudioCompletionEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*OfflineAudioCompletionEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*OfflineAudioCompletionEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*OfflineAudioCompletionEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*OfflineAudioCompletionEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*OfflineAudioCompletionEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*OfflineAudioCompletionEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*OfflineAudioCompletionEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*OfflineAudioCompletionEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*OfflineAudioCompletionEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*OfflineAudioCompletionEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*OfflineAudioCompletionEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

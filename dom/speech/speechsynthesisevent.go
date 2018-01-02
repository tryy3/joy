package speech

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*SpeechSynthesisEvent)(nil)

func New(kind string, eventinitdict *SpeechSynthesisEventInit) *SpeechSynthesisEvent {
	macro.Rewrite("new SpeechSynthesisEvent($1, $2)", kind, eventinitdict)
	return &SpeechSynthesisEvent{}
}

type SpeechSynthesisEvent struct {
}

func (*SpeechSynthesisEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*SpeechSynthesisEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*SpeechSynthesisEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*SpeechSynthesisEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*SpeechSynthesisEvent) CharIndex() (charIndex uint) {
	macro.Rewrite("$_.charIndex")
	return charIndex
}

func (*SpeechSynthesisEvent) ElapsedTime() (elapsedTime float32) {
	macro.Rewrite("$_.elapsedTime")
	return elapsedTime
}

func (*SpeechSynthesisEvent) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*SpeechSynthesisEvent) Utterance() (utterance *SpeechSynthesisUtterance) {
	macro.Rewrite("$_.utterance")
	return utterance
}

func (*SpeechSynthesisEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*SpeechSynthesisEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*SpeechSynthesisEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*SpeechSynthesisEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*SpeechSynthesisEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*SpeechSynthesisEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*SpeechSynthesisEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*SpeechSynthesisEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*SpeechSynthesisEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*SpeechSynthesisEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*SpeechSynthesisEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*SpeechSynthesisEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*SpeechSynthesisEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*SpeechSynthesisEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

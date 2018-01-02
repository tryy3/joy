package mediastreams

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*MediaStreamErrorEvent)(nil)

func NewMediaStreamErrorEvent(typearg string, eventinitdict *MediaStreamErrorEventInit) *MediaStreamErrorEvent {
	macro.Rewrite("new MediaStreamErrorEvent($1, $2)", typearg, eventinitdict)
	return &MediaStreamErrorEvent{}
}

type MediaStreamErrorEvent struct {
}

func (*MediaStreamErrorEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*MediaStreamErrorEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*MediaStreamErrorEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*MediaStreamErrorEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*MediaStreamErrorEvent) Error() (err *MediaStreamError) {
	macro.Rewrite("$_.error")
	return err
}

func (*MediaStreamErrorEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*MediaStreamErrorEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*MediaStreamErrorEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*MediaStreamErrorEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*MediaStreamErrorEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*MediaStreamErrorEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*MediaStreamErrorEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*MediaStreamErrorEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*MediaStreamErrorEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*MediaStreamErrorEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*MediaStreamErrorEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*MediaStreamErrorEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*MediaStreamErrorEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*MediaStreamErrorEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

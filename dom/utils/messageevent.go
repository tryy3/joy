package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*MessageEvent)(nil)

func NewMessageEvent(typearg string, eventinitdict *MessageEventInit) *MessageEvent {
	macro.Rewrite("new MessageEvent($1, $2)", typearg, eventinitdict)
	return &MessageEvent{}
}

type MessageEvent struct {
}

func (*MessageEvent) InitMessageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, dataArg interface{}, originArg string, lastEventIdArg string, sourceArg *window.Window) {
	macro.Rewrite("$_.initMessageEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, dataArg, originArg, lastEventIdArg, sourceArg)
}

func (*MessageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*MessageEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*MessageEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*MessageEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*MessageEvent) Data() (data interface{}) {
	macro.Rewrite("$_.data")
	return data
}

func (*MessageEvent) Origin() (origin string) {
	macro.Rewrite("$_.origin")
	return origin
}

func (*MessageEvent) Ports() (ports interface{}) {
	macro.Rewrite("$_.ports")
	return ports
}

func (*MessageEvent) Source() (source *window.Window) {
	macro.Rewrite("$_.source")
	return source
}

func (*MessageEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*MessageEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*MessageEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*MessageEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*MessageEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*MessageEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*MessageEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*MessageEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*MessageEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*MessageEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*MessageEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*MessageEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*MessageEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*MessageEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

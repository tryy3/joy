package payment

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*PaymentRequestUpdateEvent)(nil)

func New(kind string, eventinitdict *PaymentRequestUpdateEventInit) *PaymentRequestUpdateEvent {
	macro.Rewrite("new PaymentRequestUpdateEvent($1, $2)", kind, eventinitdict)
	return &PaymentRequestUpdateEvent{}
}

type PaymentRequestUpdateEvent struct {
}

func (*PaymentRequestUpdateEvent) UpdateWith(d *PaymentDetails) {
	macro.Rewrite("$_.updateWith($1)", d)
}

func (*PaymentRequestUpdateEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*PaymentRequestUpdateEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*PaymentRequestUpdateEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*PaymentRequestUpdateEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*PaymentRequestUpdateEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*PaymentRequestUpdateEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*PaymentRequestUpdateEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*PaymentRequestUpdateEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*PaymentRequestUpdateEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*PaymentRequestUpdateEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*PaymentRequestUpdateEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*PaymentRequestUpdateEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*PaymentRequestUpdateEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*PaymentRequestUpdateEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*PaymentRequestUpdateEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*PaymentRequestUpdateEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*PaymentRequestUpdateEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*PaymentRequestUpdateEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

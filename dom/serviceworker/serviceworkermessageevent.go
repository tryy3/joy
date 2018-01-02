package serviceworker

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/channelmessage"
)

var _ event.Event = (*ServiceWorkerMessageEvent)(nil)

func NewServiceWorkerMessageEvent(kind string, eventinitdict *ServiceWorkerMessageEventInit) *ServiceWorkerMessageEvent {
	macro.Rewrite("new ServiceWorkerMessageEvent($1, $2)", kind, eventinitdict)
	return &ServiceWorkerMessageEvent{}
}

type ServiceWorkerMessageEvent struct {
}

func (*ServiceWorkerMessageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*ServiceWorkerMessageEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*ServiceWorkerMessageEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*ServiceWorkerMessageEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*ServiceWorkerMessageEvent) Data() (data interface{}) {
	macro.Rewrite("$_.data")
	return data
}

func (*ServiceWorkerMessageEvent) LastEventID() (lastEventId string) {
	macro.Rewrite("$_.lastEventId")
	return lastEventId
}

func (*ServiceWorkerMessageEvent) Origin() (origin string) {
	macro.Rewrite("$_.origin")
	return origin
}

func (*ServiceWorkerMessageEvent) Ports() (ports []*channelmessage.MessagePort) {
	macro.Rewrite("$_.ports")
	return ports
}

func (*ServiceWorkerMessageEvent) Source() (source interface{}) {
	macro.Rewrite("$_.source")
	return source
}

func (*ServiceWorkerMessageEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*ServiceWorkerMessageEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*ServiceWorkerMessageEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*ServiceWorkerMessageEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*ServiceWorkerMessageEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*ServiceWorkerMessageEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*ServiceWorkerMessageEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*ServiceWorkerMessageEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*ServiceWorkerMessageEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*ServiceWorkerMessageEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*ServiceWorkerMessageEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*ServiceWorkerMessageEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*ServiceWorkerMessageEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*ServiceWorkerMessageEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

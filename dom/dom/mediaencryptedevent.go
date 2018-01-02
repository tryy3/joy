package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*MediaEncryptedEvent)(nil)

func New(kind string, eventinitdict *MediaEncryptedEventInit) *MediaEncryptedEvent {
	macro.Rewrite("new MediaEncryptedEvent($1, $2)", kind, eventinitdict)
	return &MediaEncryptedEvent{}
}

type MediaEncryptedEvent struct {
}

func (*MediaEncryptedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*MediaEncryptedEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*MediaEncryptedEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*MediaEncryptedEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*MediaEncryptedEvent) InitData() (initData []byte) {
	macro.Rewrite("$_.initData")
	return initData
}

func (*MediaEncryptedEvent) InitDataType() (initDataType string) {
	macro.Rewrite("$_.initDataType")
	return initDataType
}

func (*MediaEncryptedEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*MediaEncryptedEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*MediaEncryptedEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*MediaEncryptedEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*MediaEncryptedEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*MediaEncryptedEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*MediaEncryptedEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*MediaEncryptedEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*MediaEncryptedEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*MediaEncryptedEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*MediaEncryptedEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*MediaEncryptedEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*MediaEncryptedEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*MediaEncryptedEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ ui.UIEvent = (*FocusEvent)(nil)
var _ event.Event = (*FocusEvent)(nil)

func NewFocusEvent(typearg string, eventinitdict *FocusEventInit) *FocusEvent {
	macro.Rewrite("new FocusEvent($1, $2)", typearg, eventinitdict)
	return &FocusEvent{}
}

type FocusEvent struct {
}

func (*FocusEvent) InitFocusEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, relatedTargetArg event.EventTarget) {
	macro.Rewrite("$_.initFocusEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, relatedTargetArg)
}

func (*FocusEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*FocusEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*FocusEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*FocusEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*FocusEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*FocusEvent) RelatedTarget() (relatedTarget event.EventTarget) {
	macro.Rewrite("$_.relatedTarget")
	return relatedTarget
}

func (*FocusEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*FocusEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*FocusEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*FocusEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*FocusEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*FocusEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*FocusEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*FocusEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*FocusEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*FocusEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*FocusEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*FocusEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*FocusEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*FocusEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*FocusEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*FocusEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

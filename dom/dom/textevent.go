package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ ui.UIEvent = (*TextEvent)(nil)
var _ event.Event = (*TextEvent)(nil)

type TextEvent struct {
}

func (*TextEvent) InitTextEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, dataArg string, inputMethod uint, locale string) {
	macro.Rewrite("$_.initTextEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, dataArg, inputMethod, locale)
}

func (*TextEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*TextEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*TextEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*TextEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*TextEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*TextEvent) Data() (data string) {
	macro.Rewrite("$_.data")
	return data
}

func (*TextEvent) InputMethod() (inputMethod uint) {
	macro.Rewrite("$_.inputMethod")
	return inputMethod
}

func (*TextEvent) Locale() (locale string) {
	macro.Rewrite("$_.locale")
	return locale
}

func (*TextEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*TextEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*TextEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*TextEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*TextEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*TextEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*TextEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*TextEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*TextEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*TextEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*TextEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*TextEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*TextEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*TextEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*TextEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*TextEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

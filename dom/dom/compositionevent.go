package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ ui.UIEvent = (*CompositionEvent)(nil)
var _ event.Event = (*CompositionEvent)(nil)

func New(typearg string, eventinitdict *CompositionEventInit) *CompositionEvent {
	macro.Rewrite("new CompositionEvent($1, $2)", typearg, eventinitdict)
	return &CompositionEvent{}
}

type CompositionEvent struct {
}

func (*CompositionEvent) InitCompositionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, dataArg string, locale string) {
	macro.Rewrite("$_.initCompositionEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, viewArg, dataArg, locale)
}

func (*CompositionEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*CompositionEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*CompositionEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*CompositionEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*CompositionEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*CompositionEvent) Data() (data string) {
	macro.Rewrite("$_.data")
	return data
}

func (*CompositionEvent) Locale() (locale string) {
	macro.Rewrite("$_.locale")
	return locale
}

func (*CompositionEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*CompositionEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*CompositionEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*CompositionEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*CompositionEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*CompositionEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*CompositionEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*CompositionEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*CompositionEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*CompositionEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*CompositionEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*CompositionEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*CompositionEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*CompositionEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*CompositionEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*CompositionEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

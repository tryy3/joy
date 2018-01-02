package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ ui.UIEvent = (*OverflowEvent)(nil)
var _ event.Event = (*OverflowEvent)(nil)

type OverflowEvent struct {
}

func (*OverflowEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*OverflowEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*OverflowEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*OverflowEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*OverflowEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*OverflowEvent) HorizontalOverflow() (horizontalOverflow bool) {
	macro.Rewrite("$_.horizontalOverflow")
	return horizontalOverflow
}

func (*OverflowEvent) Orient() (orient uint) {
	macro.Rewrite("$_.orient")
	return orient
}

func (*OverflowEvent) VerticalOverflow() (verticalOverflow bool) {
	macro.Rewrite("$_.verticalOverflow")
	return verticalOverflow
}

func (*OverflowEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*OverflowEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*OverflowEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*OverflowEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*OverflowEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*OverflowEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*OverflowEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*OverflowEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*OverflowEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*OverflowEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*OverflowEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*OverflowEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*OverflowEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*OverflowEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*OverflowEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*OverflowEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

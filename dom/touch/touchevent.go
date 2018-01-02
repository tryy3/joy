package touch

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ ui.UIEvent = (*TouchEvent)(nil)
var _ event.Event = (*TouchEvent)(nil)

type TouchEvent struct {
}

func (*TouchEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*TouchEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*TouchEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*TouchEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*TouchEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*TouchEvent) AltKey() (altKey bool) {
	macro.Rewrite("$_.altKey")
	return altKey
}

func (*TouchEvent) ChangedTouches() (changedTouches *TouchList) {
	macro.Rewrite("$_.changedTouches")
	return changedTouches
}

func (*TouchEvent) CharCode() (charCode int8) {
	macro.Rewrite("$_.charCode")
	return charCode
}

func (*TouchEvent) CtrlKey() (ctrlKey bool) {
	macro.Rewrite("$_.ctrlKey")
	return ctrlKey
}

func (*TouchEvent) KeyCode() (keyCode int8) {
	macro.Rewrite("$_.keyCode")
	return keyCode
}

func (*TouchEvent) MetaKey() (metaKey bool) {
	macro.Rewrite("$_.metaKey")
	return metaKey
}

func (*TouchEvent) ShiftKey() (shiftKey bool) {
	macro.Rewrite("$_.shiftKey")
	return shiftKey
}

func (*TouchEvent) TargetTouches() (targetTouches *TouchList) {
	macro.Rewrite("$_.targetTouches")
	return targetTouches
}

func (*TouchEvent) Touches() (touches *TouchList) {
	macro.Rewrite("$_.touches")
	return touches
}

func (*TouchEvent) Which() (which uint8) {
	macro.Rewrite("$_.which")
	return which
}

func (*TouchEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*TouchEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*TouchEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*TouchEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*TouchEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*TouchEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*TouchEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*TouchEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*TouchEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*TouchEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*TouchEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*TouchEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*TouchEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*TouchEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*TouchEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*TouchEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

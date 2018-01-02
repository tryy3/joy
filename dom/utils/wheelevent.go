package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/mouse"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/window"
)

var _ mouse.MouseEvent = (*WheelEvent)(nil)
var _ ui.UIEvent = (*WheelEvent)(nil)
var _ event.Event = (*WheelEvent)(nil)

func NewWheelEvent(typearg string, eventinitdict *WheelEventInit) *WheelEvent {
	macro.Rewrite("new WheelEvent($1, $2)", typearg, eventinitdict)
	return &WheelEvent{}
}

type WheelEvent struct {
}

func (*WheelEvent) GetCurrentPoint(element element.Element) {
	macro.Rewrite("$_.getCurrentPoint($1)", element)
}

func (*WheelEvent) InitWheelEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, buttonArg uint8, relatedTargetArg event.EventTarget, modifiersListArg string, deltaXArg int, deltaYArg int, deltaZArg int, deltaMode uint) {
	macro.Rewrite("$_.initWheelEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, buttonArg, relatedTargetArg, modifiersListArg, deltaXArg, deltaYArg, deltaZArg, deltaMode)
}

func (*WheelEvent) GetModifierState(keyArg string) (b bool) {
	macro.Rewrite("$_.getModifierState($1)", keyArg)
	return b
}

func (*WheelEvent) InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg event.EventTarget) {
	macro.Rewrite("$_.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

func (*WheelEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*WheelEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*WheelEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*WheelEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*WheelEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*WheelEvent) DeltaMode() (deltaMode uint) {
	macro.Rewrite("$_.deltaMode")
	return deltaMode
}

func (*WheelEvent) DeltaX() (deltaX int) {
	macro.Rewrite("$_.deltaX")
	return deltaX
}

func (*WheelEvent) DeltaY() (deltaY int) {
	macro.Rewrite("$_.deltaY")
	return deltaY
}

func (*WheelEvent) DeltaZ() (deltaZ int) {
	macro.Rewrite("$_.deltaZ")
	return deltaZ
}

func (*WheelEvent) WheelDelta() (wheelDelta int) {
	macro.Rewrite("$_.wheelDelta")
	return wheelDelta
}

func (*WheelEvent) WheelDeltaX() (wheelDeltaX int) {
	macro.Rewrite("$_.wheelDeltaX")
	return wheelDeltaX
}

func (*WheelEvent) WheelDeltaY() (wheelDeltaY int) {
	macro.Rewrite("$_.wheelDeltaY")
	return wheelDeltaY
}

func (*WheelEvent) AltKey() (altKey bool) {
	macro.Rewrite("$_.altKey")
	return altKey
}

func (*WheelEvent) Button() (button uint8) {
	macro.Rewrite("$_.button")
	return button
}

func (*WheelEvent) Buttons() (buttons uint8) {
	macro.Rewrite("$_.buttons")
	return buttons
}

func (*WheelEvent) ClientX() (clientX int) {
	macro.Rewrite("$_.clientX")
	return clientX
}

func (*WheelEvent) ClientY() (clientY int) {
	macro.Rewrite("$_.clientY")
	return clientY
}

func (*WheelEvent) CtrlKey() (ctrlKey bool) {
	macro.Rewrite("$_.ctrlKey")
	return ctrlKey
}

func (*WheelEvent) FromElement() (fromElement element.Element) {
	macro.Rewrite("$_.fromElement")
	return fromElement
}

func (*WheelEvent) LayerX() (layerX int) {
	macro.Rewrite("$_.layerX")
	return layerX
}

func (*WheelEvent) LayerY() (layerY int) {
	macro.Rewrite("$_.layerY")
	return layerY
}

func (*WheelEvent) MetaKey() (metaKey bool) {
	macro.Rewrite("$_.metaKey")
	return metaKey
}

func (*WheelEvent) MovementX() (movementX int) {
	macro.Rewrite("$_.movementX")
	return movementX
}

func (*WheelEvent) MovementY() (movementY int) {
	macro.Rewrite("$_.movementY")
	return movementY
}

func (*WheelEvent) OffsetX() (offsetX int) {
	macro.Rewrite("$_.offsetX")
	return offsetX
}

func (*WheelEvent) OffsetY() (offsetY int) {
	macro.Rewrite("$_.offsetY")
	return offsetY
}

func (*WheelEvent) PageX() (pageX int) {
	macro.Rewrite("$_.pageX")
	return pageX
}

func (*WheelEvent) PageY() (pageY int) {
	macro.Rewrite("$_.pageY")
	return pageY
}

func (*WheelEvent) RelatedTarget() (relatedTarget event.EventTarget) {
	macro.Rewrite("$_.relatedTarget")
	return relatedTarget
}

func (*WheelEvent) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

func (*WheelEvent) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

func (*WheelEvent) ShiftKey() (shiftKey bool) {
	macro.Rewrite("$_.shiftKey")
	return shiftKey
}

func (*WheelEvent) ToElement() (toElement element.Element) {
	macro.Rewrite("$_.toElement")
	return toElement
}

func (*WheelEvent) Which() (which uint8) {
	macro.Rewrite("$_.which")
	return which
}

func (*WheelEvent) X() (x int) {
	macro.Rewrite("$_.x")
	return x
}

func (*WheelEvent) Y() (y int) {
	macro.Rewrite("$_.y")
	return y
}

func (*WheelEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*WheelEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*WheelEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*WheelEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*WheelEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*WheelEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*WheelEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*WheelEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*WheelEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*WheelEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*WheelEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*WheelEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*WheelEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*WheelEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*WheelEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*WheelEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

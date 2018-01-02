package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/mouse"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/window"
)

var _ mouse.MouseEvent = (*PointerEvent)(nil)
var _ ui.UIEvent = (*PointerEvent)(nil)
var _ event.Event = (*PointerEvent)(nil)

func NewPointerEvent(typearg string, eventinitdict *PointerEventInit) *PointerEvent {
	macro.Rewrite("new PointerEvent($1, $2)", typearg, eventinitdict)
	return &PointerEvent{}
}

type PointerEvent struct {
}

func (*PointerEvent) GetCurrentPoint(element element.Element) {
	macro.Rewrite("$_.getCurrentPoint($1)", element)
}

func (*PointerEvent) GetIntermediatePoints(element element.Element) {
	macro.Rewrite("$_.getIntermediatePoints($1)", element)
}

func (*PointerEvent) InitPointerEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg event.EventTarget, offsetXArg float32, offsetYArg float32, widthArg int, heightArg int, pressure float32, rotation int, tiltX int, tiltY int, pointerIdArg int, pointerType interface{}, hwTimestampArg uint64, isPrimary bool) {
	macro.Rewrite("$_.initPointerEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, offsetXArg, offsetYArg, widthArg, heightArg, pressure, rotation, tiltX, tiltY, pointerIdArg, pointerType, hwTimestampArg, isPrimary)
}

func (*PointerEvent) GetModifierState(keyArg string) (b bool) {
	macro.Rewrite("$_.getModifierState($1)", keyArg)
	return b
}

func (*PointerEvent) InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg event.EventTarget) {
	macro.Rewrite("$_.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

func (*PointerEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*PointerEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*PointerEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*PointerEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*PointerEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*PointerEvent) CurrentPoint() (currentPoint interface{}) {
	macro.Rewrite("$_.currentPoint")
	return currentPoint
}

func (*PointerEvent) Height() (height int) {
	macro.Rewrite("$_.height")
	return height
}

func (*PointerEvent) HwTimestamp() (hwTimestamp uint64) {
	macro.Rewrite("$_.hwTimestamp")
	return hwTimestamp
}

func (*PointerEvent) IntermediatePoints() (intermediatePoints interface{}) {
	macro.Rewrite("$_.intermediatePoints")
	return intermediatePoints
}

func (*PointerEvent) IsPrimary() (isPrimary bool) {
	macro.Rewrite("$_.isPrimary")
	return isPrimary
}

func (*PointerEvent) PointerID() (pointerId int) {
	macro.Rewrite("$_.pointerId")
	return pointerId
}

func (*PointerEvent) PointerType() (pointerType interface{}) {
	macro.Rewrite("$_.pointerType")
	return pointerType
}

func (*PointerEvent) Pressure() (pressure float32) {
	macro.Rewrite("$_.pressure")
	return pressure
}

func (*PointerEvent) Rotation() (rotation int) {
	macro.Rewrite("$_.rotation")
	return rotation
}

func (*PointerEvent) TiltX() (tiltX int) {
	macro.Rewrite("$_.tiltX")
	return tiltX
}

func (*PointerEvent) TiltY() (tiltY int) {
	macro.Rewrite("$_.tiltY")
	return tiltY
}

func (*PointerEvent) Width() (width int) {
	macro.Rewrite("$_.width")
	return width
}

func (*PointerEvent) AltKey() (altKey bool) {
	macro.Rewrite("$_.altKey")
	return altKey
}

func (*PointerEvent) Button() (button uint8) {
	macro.Rewrite("$_.button")
	return button
}

func (*PointerEvent) Buttons() (buttons uint8) {
	macro.Rewrite("$_.buttons")
	return buttons
}

func (*PointerEvent) ClientX() (clientX int) {
	macro.Rewrite("$_.clientX")
	return clientX
}

func (*PointerEvent) ClientY() (clientY int) {
	macro.Rewrite("$_.clientY")
	return clientY
}

func (*PointerEvent) CtrlKey() (ctrlKey bool) {
	macro.Rewrite("$_.ctrlKey")
	return ctrlKey
}

func (*PointerEvent) FromElement() (fromElement element.Element) {
	macro.Rewrite("$_.fromElement")
	return fromElement
}

func (*PointerEvent) LayerX() (layerX int) {
	macro.Rewrite("$_.layerX")
	return layerX
}

func (*PointerEvent) LayerY() (layerY int) {
	macro.Rewrite("$_.layerY")
	return layerY
}

func (*PointerEvent) MetaKey() (metaKey bool) {
	macro.Rewrite("$_.metaKey")
	return metaKey
}

func (*PointerEvent) MovementX() (movementX int) {
	macro.Rewrite("$_.movementX")
	return movementX
}

func (*PointerEvent) MovementY() (movementY int) {
	macro.Rewrite("$_.movementY")
	return movementY
}

func (*PointerEvent) OffsetX() (offsetX int) {
	macro.Rewrite("$_.offsetX")
	return offsetX
}

func (*PointerEvent) OffsetY() (offsetY int) {
	macro.Rewrite("$_.offsetY")
	return offsetY
}

func (*PointerEvent) PageX() (pageX int) {
	macro.Rewrite("$_.pageX")
	return pageX
}

func (*PointerEvent) PageY() (pageY int) {
	macro.Rewrite("$_.pageY")
	return pageY
}

func (*PointerEvent) RelatedTarget() (relatedTarget event.EventTarget) {
	macro.Rewrite("$_.relatedTarget")
	return relatedTarget
}

func (*PointerEvent) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

func (*PointerEvent) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

func (*PointerEvent) ShiftKey() (shiftKey bool) {
	macro.Rewrite("$_.shiftKey")
	return shiftKey
}

func (*PointerEvent) ToElement() (toElement element.Element) {
	macro.Rewrite("$_.toElement")
	return toElement
}

func (*PointerEvent) Which() (which uint8) {
	macro.Rewrite("$_.which")
	return which
}

func (*PointerEvent) X() (x int) {
	macro.Rewrite("$_.x")
	return x
}

func (*PointerEvent) Y() (y int) {
	macro.Rewrite("$_.y")
	return y
}

func (*PointerEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*PointerEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*PointerEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*PointerEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*PointerEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*PointerEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*PointerEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*PointerEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*PointerEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*PointerEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*PointerEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*PointerEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*PointerEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*PointerEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*PointerEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*PointerEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

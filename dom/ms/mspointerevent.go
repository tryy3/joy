package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/mouse"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/window"
)

var _ mouse.MouseEvent = (*MSPointerEvent)(nil)
var _ ui.UIEvent = (*MSPointerEvent)(nil)
var _ event.Event = (*MSPointerEvent)(nil)

func NewMSPointerEvent(typearg string, eventinitdict *utils.PointerEventInit) *MSPointerEvent {
	macro.Rewrite("new MSPointerEvent($1, $2)", typearg, eventinitdict)
	return &MSPointerEvent{}
}

type MSPointerEvent struct {
}

func (*MSPointerEvent) GetCurrentPoint(element element.Element) {
	macro.Rewrite("$_.getCurrentPoint($1)", element)
}

func (*MSPointerEvent) GetIntermediatePoints(element element.Element) {
	macro.Rewrite("$_.getIntermediatePoints($1)", element)
}

func (*MSPointerEvent) InitPointerEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg event.EventTarget, offsetXArg float32, offsetYArg float32, widthArg int, heightArg int, pressure float32, rotation int, tiltX int, tiltY int, pointerIdArg int, pointerType interface{}, hwTimestampArg uint64, isPrimary bool) {
	macro.Rewrite("$_.initPointerEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, offsetXArg, offsetYArg, widthArg, heightArg, pressure, rotation, tiltX, tiltY, pointerIdArg, pointerType, hwTimestampArg, isPrimary)
}

func (*MSPointerEvent) GetModifierState(keyArg string) (b bool) {
	macro.Rewrite("$_.getModifierState($1)", keyArg)
	return b
}

func (*MSPointerEvent) InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg event.EventTarget) {
	macro.Rewrite("$_.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

func (*MSPointerEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*MSPointerEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*MSPointerEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*MSPointerEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*MSPointerEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*MSPointerEvent) CurrentPoint() (currentPoint interface{}) {
	macro.Rewrite("$_.currentPoint")
	return currentPoint
}

func (*MSPointerEvent) Height() (height int) {
	macro.Rewrite("$_.height")
	return height
}

func (*MSPointerEvent) HwTimestamp() (hwTimestamp uint64) {
	macro.Rewrite("$_.hwTimestamp")
	return hwTimestamp
}

func (*MSPointerEvent) IntermediatePoints() (intermediatePoints interface{}) {
	macro.Rewrite("$_.intermediatePoints")
	return intermediatePoints
}

func (*MSPointerEvent) IsPrimary() (isPrimary bool) {
	macro.Rewrite("$_.isPrimary")
	return isPrimary
}

func (*MSPointerEvent) PointerID() (pointerId int) {
	macro.Rewrite("$_.pointerId")
	return pointerId
}

func (*MSPointerEvent) PointerType() (pointerType interface{}) {
	macro.Rewrite("$_.pointerType")
	return pointerType
}

func (*MSPointerEvent) Pressure() (pressure float32) {
	macro.Rewrite("$_.pressure")
	return pressure
}

func (*MSPointerEvent) Rotation() (rotation int) {
	macro.Rewrite("$_.rotation")
	return rotation
}

func (*MSPointerEvent) TiltX() (tiltX int) {
	macro.Rewrite("$_.tiltX")
	return tiltX
}

func (*MSPointerEvent) TiltY() (tiltY int) {
	macro.Rewrite("$_.tiltY")
	return tiltY
}

func (*MSPointerEvent) Width() (width int) {
	macro.Rewrite("$_.width")
	return width
}

func (*MSPointerEvent) AltKey() (altKey bool) {
	macro.Rewrite("$_.altKey")
	return altKey
}

func (*MSPointerEvent) Button() (button uint8) {
	macro.Rewrite("$_.button")
	return button
}

func (*MSPointerEvent) Buttons() (buttons uint8) {
	macro.Rewrite("$_.buttons")
	return buttons
}

func (*MSPointerEvent) ClientX() (clientX int) {
	macro.Rewrite("$_.clientX")
	return clientX
}

func (*MSPointerEvent) ClientY() (clientY int) {
	macro.Rewrite("$_.clientY")
	return clientY
}

func (*MSPointerEvent) CtrlKey() (ctrlKey bool) {
	macro.Rewrite("$_.ctrlKey")
	return ctrlKey
}

func (*MSPointerEvent) FromElement() (fromElement element.Element) {
	macro.Rewrite("$_.fromElement")
	return fromElement
}

func (*MSPointerEvent) LayerX() (layerX int) {
	macro.Rewrite("$_.layerX")
	return layerX
}

func (*MSPointerEvent) LayerY() (layerY int) {
	macro.Rewrite("$_.layerY")
	return layerY
}

func (*MSPointerEvent) MetaKey() (metaKey bool) {
	macro.Rewrite("$_.metaKey")
	return metaKey
}

func (*MSPointerEvent) MovementX() (movementX int) {
	macro.Rewrite("$_.movementX")
	return movementX
}

func (*MSPointerEvent) MovementY() (movementY int) {
	macro.Rewrite("$_.movementY")
	return movementY
}

func (*MSPointerEvent) OffsetX() (offsetX int) {
	macro.Rewrite("$_.offsetX")
	return offsetX
}

func (*MSPointerEvent) OffsetY() (offsetY int) {
	macro.Rewrite("$_.offsetY")
	return offsetY
}

func (*MSPointerEvent) PageX() (pageX int) {
	macro.Rewrite("$_.pageX")
	return pageX
}

func (*MSPointerEvent) PageY() (pageY int) {
	macro.Rewrite("$_.pageY")
	return pageY
}

func (*MSPointerEvent) RelatedTarget() (relatedTarget event.EventTarget) {
	macro.Rewrite("$_.relatedTarget")
	return relatedTarget
}

func (*MSPointerEvent) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

func (*MSPointerEvent) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

func (*MSPointerEvent) ShiftKey() (shiftKey bool) {
	macro.Rewrite("$_.shiftKey")
	return shiftKey
}

func (*MSPointerEvent) ToElement() (toElement element.Element) {
	macro.Rewrite("$_.toElement")
	return toElement
}

func (*MSPointerEvent) Which() (which uint8) {
	macro.Rewrite("$_.which")
	return which
}

func (*MSPointerEvent) X() (x int) {
	macro.Rewrite("$_.x")
	return x
}

func (*MSPointerEvent) Y() (y int) {
	macro.Rewrite("$_.y")
	return y
}

func (*MSPointerEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*MSPointerEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*MSPointerEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*MSPointerEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*MSPointerEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*MSPointerEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*MSPointerEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*MSPointerEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*MSPointerEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*MSPointerEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*MSPointerEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*MSPointerEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*MSPointerEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*MSPointerEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*MSPointerEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*MSPointerEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

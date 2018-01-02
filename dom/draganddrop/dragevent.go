package draganddrop

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/mouse"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/dom/element"
)

var _ mouse.MouseEvent = (*DragEvent)(nil)
var _ ui.UIEvent = (*DragEvent)(nil)
var _ event.Event = (*DragEvent)(nil)

type DragEvent struct {
}

func (*DragEvent) InitDragEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg event.EventTarget, dataTransferArg *DataTransfer) {
	macro.Rewrite("$_.initDragEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, dataTransferArg)
}

func (*DragEvent) MsConvertURL(file *file.File, targetType string, targetURL *string) {
	macro.Rewrite("$_.msConvertURL($1, $2, $3)", file, targetType, targetURL)
}

func (*DragEvent) GetModifierState(keyArg string) (b bool) {
	macro.Rewrite("$_.getModifierState($1)", keyArg)
	return b
}

func (*DragEvent) InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg event.EventTarget) {
	macro.Rewrite("$_.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

func (*DragEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*DragEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*DragEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*DragEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*DragEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*DragEvent) DataTransfer() (dataTransfer *DataTransfer) {
	macro.Rewrite("$_.dataTransfer")
	return dataTransfer
}

func (*DragEvent) AltKey() (altKey bool) {
	macro.Rewrite("$_.altKey")
	return altKey
}

func (*DragEvent) Button() (button uint8) {
	macro.Rewrite("$_.button")
	return button
}

func (*DragEvent) Buttons() (buttons uint8) {
	macro.Rewrite("$_.buttons")
	return buttons
}

func (*DragEvent) ClientX() (clientX int) {
	macro.Rewrite("$_.clientX")
	return clientX
}

func (*DragEvent) ClientY() (clientY int) {
	macro.Rewrite("$_.clientY")
	return clientY
}

func (*DragEvent) CtrlKey() (ctrlKey bool) {
	macro.Rewrite("$_.ctrlKey")
	return ctrlKey
}

func (*DragEvent) FromElement() (fromElement element.Element) {
	macro.Rewrite("$_.fromElement")
	return fromElement
}

func (*DragEvent) LayerX() (layerX int) {
	macro.Rewrite("$_.layerX")
	return layerX
}

func (*DragEvent) LayerY() (layerY int) {
	macro.Rewrite("$_.layerY")
	return layerY
}

func (*DragEvent) MetaKey() (metaKey bool) {
	macro.Rewrite("$_.metaKey")
	return metaKey
}

func (*DragEvent) MovementX() (movementX int) {
	macro.Rewrite("$_.movementX")
	return movementX
}

func (*DragEvent) MovementY() (movementY int) {
	macro.Rewrite("$_.movementY")
	return movementY
}

func (*DragEvent) OffsetX() (offsetX int) {
	macro.Rewrite("$_.offsetX")
	return offsetX
}

func (*DragEvent) OffsetY() (offsetY int) {
	macro.Rewrite("$_.offsetY")
	return offsetY
}

func (*DragEvent) PageX() (pageX int) {
	macro.Rewrite("$_.pageX")
	return pageX
}

func (*DragEvent) PageY() (pageY int) {
	macro.Rewrite("$_.pageY")
	return pageY
}

func (*DragEvent) RelatedTarget() (relatedTarget event.EventTarget) {
	macro.Rewrite("$_.relatedTarget")
	return relatedTarget
}

func (*DragEvent) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

func (*DragEvent) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

func (*DragEvent) ShiftKey() (shiftKey bool) {
	macro.Rewrite("$_.shiftKey")
	return shiftKey
}

func (*DragEvent) ToElement() (toElement element.Element) {
	macro.Rewrite("$_.toElement")
	return toElement
}

func (*DragEvent) Which() (which uint8) {
	macro.Rewrite("$_.which")
	return which
}

func (*DragEvent) X() (x int) {
	macro.Rewrite("$_.x")
	return x
}

func (*DragEvent) Y() (y int) {
	macro.Rewrite("$_.y")
	return y
}

func (*DragEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*DragEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*DragEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*DragEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*DragEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*DragEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*DragEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*DragEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*DragEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*DragEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*DragEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*DragEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*DragEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*DragEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*DragEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*DragEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

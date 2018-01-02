package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ ui.UIEvent = (*MSGestureEvent)(nil)
var _ event.Event = (*MSGestureEvent)(nil)

type MSGestureEvent struct {
}

func (*MSGestureEvent) InitGestureEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, offsetXArg float32, offsetYArg float32, translationXArg float32, translationYArg float32, scaleArg float32, expansionArg float32, rotationArg float32, velocityXArg float32, velocityYArg float32, velocityExpansionArg float32, velocityAngularArg float32, hwTimestampArg uint64) {
	macro.Rewrite("$_.initGestureEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, offsetXArg, offsetYArg, translationXArg, translationYArg, scaleArg, expansionArg, rotationArg, velocityXArg, velocityYArg, velocityExpansionArg, velocityAngularArg, hwTimestampArg)
}

func (*MSGestureEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*MSGestureEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*MSGestureEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*MSGestureEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*MSGestureEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*MSGestureEvent) ClientX() (clientX float32) {
	macro.Rewrite("$_.clientX")
	return clientX
}

func (*MSGestureEvent) ClientY() (clientY float32) {
	macro.Rewrite("$_.clientY")
	return clientY
}

func (*MSGestureEvent) Expansion() (expansion float32) {
	macro.Rewrite("$_.expansion")
	return expansion
}

func (*MSGestureEvent) GestureObject() (gestureObject interface{}) {
	macro.Rewrite("$_.gestureObject")
	return gestureObject
}

func (*MSGestureEvent) HwTimestamp() (hwTimestamp uint64) {
	macro.Rewrite("$_.hwTimestamp")
	return hwTimestamp
}

func (*MSGestureEvent) OffsetX() (offsetX float32) {
	macro.Rewrite("$_.offsetX")
	return offsetX
}

func (*MSGestureEvent) OffsetY() (offsetY float32) {
	macro.Rewrite("$_.offsetY")
	return offsetY
}

func (*MSGestureEvent) Rotation() (rotation float32) {
	macro.Rewrite("$_.rotation")
	return rotation
}

func (*MSGestureEvent) Scale() (scale float32) {
	macro.Rewrite("$_.scale")
	return scale
}

func (*MSGestureEvent) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

func (*MSGestureEvent) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

func (*MSGestureEvent) TranslationX() (translationX float32) {
	macro.Rewrite("$_.translationX")
	return translationX
}

func (*MSGestureEvent) TranslationY() (translationY float32) {
	macro.Rewrite("$_.translationY")
	return translationY
}

func (*MSGestureEvent) VelocityAngular() (velocityAngular float32) {
	macro.Rewrite("$_.velocityAngular")
	return velocityAngular
}

func (*MSGestureEvent) VelocityExpansion() (velocityExpansion float32) {
	macro.Rewrite("$_.velocityExpansion")
	return velocityExpansion
}

func (*MSGestureEvent) VelocityX() (velocityX float32) {
	macro.Rewrite("$_.velocityX")
	return velocityX
}

func (*MSGestureEvent) VelocityY() (velocityY float32) {
	macro.Rewrite("$_.velocityY")
	return velocityY
}

func (*MSGestureEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*MSGestureEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*MSGestureEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*MSGestureEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*MSGestureEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*MSGestureEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*MSGestureEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*MSGestureEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*MSGestureEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*MSGestureEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*MSGestureEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*MSGestureEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*MSGestureEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*MSGestureEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*MSGestureEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*MSGestureEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

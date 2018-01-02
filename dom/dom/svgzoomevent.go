package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ ui.UIEvent = (*SVGZoomEvent)(nil)
var _ event.Event = (*SVGZoomEvent)(nil)

type SVGZoomEvent struct {
}

func (*SVGZoomEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*SVGZoomEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*SVGZoomEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*SVGZoomEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*SVGZoomEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*SVGZoomEvent) NewScale() (newScale float32) {
	macro.Rewrite("$_.newScale")
	return newScale
}

func (*SVGZoomEvent) NewTranslate() (newTranslate *SVGPoint) {
	macro.Rewrite("$_.newTranslate")
	return newTranslate
}

func (*SVGZoomEvent) PreviousScale() (previousScale float32) {
	macro.Rewrite("$_.previousScale")
	return previousScale
}

func (*SVGZoomEvent) PreviousTranslate() (previousTranslate *SVGPoint) {
	macro.Rewrite("$_.previousTranslate")
	return previousTranslate
}

func (*SVGZoomEvent) ZoomRectScreen() (zoomRectScreen *SVGRect) {
	macro.Rewrite("$_.zoomRectScreen")
	return zoomRectScreen
}

func (*SVGZoomEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*SVGZoomEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*SVGZoomEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*SVGZoomEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*SVGZoomEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*SVGZoomEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*SVGZoomEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*SVGZoomEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*SVGZoomEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*SVGZoomEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*SVGZoomEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*SVGZoomEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*SVGZoomEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*SVGZoomEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*SVGZoomEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*SVGZoomEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

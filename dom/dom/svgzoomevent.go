package dom


import (
	"github.com/matthewmueller/joy/dom/svgpoint"
	"github.com/matthewmueller/joy/dom/svgrect"
	"github.com/matthewmueller/joy/macro"
)

var _ UIEvent = (*SVGZoomEvent)(nil)
var _ Event = (*SVGZoomEvent)(nil)

// SVGZoomEvent struct
// js:"SVGZoomEvent,omit"
type SVGZoomEvent struct {
}

// InitUIEvent fn
// js:"initUIEvent"
func (*SVGZoomEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*SVGZoomEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*SVGZoomEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*SVGZoomEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*SVGZoomEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// NewScale prop
// js:"newScale"
func (*SVGZoomEvent) NewScale() (newScale float32) {
	macro.Rewrite("$_.newScale")
	return newScale
}

// NewTranslate prop
// js:"newTranslate"
func (*SVGZoomEvent) NewTranslate() (newTranslate *svgpoint.SVGPoint) {
	macro.Rewrite("$_.newTranslate")
	return newTranslate
}

// PreviousScale prop
// js:"previousScale"
func (*SVGZoomEvent) PreviousScale() (previousScale float32) {
	macro.Rewrite("$_.previousScale")
	return previousScale
}

// PreviousTranslate prop
// js:"previousTranslate"
func (*SVGZoomEvent) PreviousTranslate() (previousTranslate *svgpoint.SVGPoint) {
	macro.Rewrite("$_.previousTranslate")
	return previousTranslate
}

// ZoomRectScreen prop
// js:"zoomRectScreen"
func (*SVGZoomEvent) ZoomRectScreen() (zoomRectScreen *svgrect.SVGRect) {
	macro.Rewrite("$_.zoomRectScreen")
	return zoomRectScreen
}

// Detail prop
// js:"detail"
func (*SVGZoomEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*SVGZoomEvent) View() (view *Window) {
	macro.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*SVGZoomEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*SVGZoomEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*SVGZoomEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*SVGZoomEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*SVGZoomEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*SVGZoomEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*SVGZoomEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*SVGZoomEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*SVGZoomEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*SVGZoomEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*SVGZoomEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*SVGZoomEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*SVGZoomEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*SVGZoomEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

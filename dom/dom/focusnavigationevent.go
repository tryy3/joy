package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/navigation"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*FocusNavigationEvent)(nil)

func New(kind string, eventinitdict *FocusNavigationEventInit) *FocusNavigationEvent {
	macro.Rewrite("new FocusNavigationEvent($1, $2)", kind, eventinitdict)
	return &FocusNavigationEvent{}
}

type FocusNavigationEvent struct {
}

func (*FocusNavigationEvent) RequestFocus() {
	macro.Rewrite("$_.requestFocus()")
}

func (*FocusNavigationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*FocusNavigationEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*FocusNavigationEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*FocusNavigationEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*FocusNavigationEvent) NavigationReason() (navigationReason *navigation.NavigationReason) {
	macro.Rewrite("$_.navigationReason")
	return navigationReason
}

func (*FocusNavigationEvent) OriginHeight() (originHeight float32) {
	macro.Rewrite("$_.originHeight")
	return originHeight
}

func (*FocusNavigationEvent) OriginLeft() (originLeft float32) {
	macro.Rewrite("$_.originLeft")
	return originLeft
}

func (*FocusNavigationEvent) OriginTop() (originTop float32) {
	macro.Rewrite("$_.originTop")
	return originTop
}

func (*FocusNavigationEvent) OriginWidth() (originWidth float32) {
	macro.Rewrite("$_.originWidth")
	return originWidth
}

func (*FocusNavigationEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*FocusNavigationEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*FocusNavigationEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*FocusNavigationEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*FocusNavigationEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*FocusNavigationEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*FocusNavigationEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*FocusNavigationEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*FocusNavigationEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*FocusNavigationEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*FocusNavigationEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*FocusNavigationEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*FocusNavigationEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*FocusNavigationEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

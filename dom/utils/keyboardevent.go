package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

var _ ui.UIEvent = (*KeyboardEvent)(nil)
var _ event.Event = (*KeyboardEvent)(nil)

func NewKeyboardEvent(typearg string, eventinitdict *KeyboardEventInit) *KeyboardEvent {
	macro.Rewrite("new KeyboardEvent($1, $2)", typearg, eventinitdict)
	return &KeyboardEvent{}
}

type KeyboardEvent struct {
}

func (*KeyboardEvent) GetModifierState(keyArg string) (b bool) {
	macro.Rewrite("$_.getModifierState($1)", keyArg)
	return b
}

func (*KeyboardEvent) InitKeyboardEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, keyArg string, locationArg uint, modifiersListArg string, repeat bool, locale string) {
	macro.Rewrite("$_.initKeyboardEvent($1, $2, $3, $4, $5, $6, $7, $8, $9)", typeArg, canBubbleArg, cancelableArg, viewArg, keyArg, locationArg, modifiersListArg, repeat, locale)
}

func (*KeyboardEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*KeyboardEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*KeyboardEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*KeyboardEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*KeyboardEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*KeyboardEvent) AltKey() (altKey bool) {
	macro.Rewrite("$_.altKey")
	return altKey
}

func (*KeyboardEvent) Char() (char string) {
	macro.Rewrite("$_.char")
	return char
}

func (*KeyboardEvent) CharCode() (charCode int8) {
	macro.Rewrite("$_.charCode")
	return charCode
}

func (*KeyboardEvent) CtrlKey() (ctrlKey bool) {
	macro.Rewrite("$_.ctrlKey")
	return ctrlKey
}

func (*KeyboardEvent) Key() (key string) {
	macro.Rewrite("$_.key")
	return key
}

func (*KeyboardEvent) KeyCode() (keyCode int8) {
	macro.Rewrite("$_.keyCode")
	return keyCode
}

func (*KeyboardEvent) Locale() (locale string) {
	macro.Rewrite("$_.locale")
	return locale
}

func (*KeyboardEvent) Location() (location uint) {
	macro.Rewrite("$_.location")
	return location
}

func (*KeyboardEvent) MetaKey() (metaKey bool) {
	macro.Rewrite("$_.metaKey")
	return metaKey
}

func (*KeyboardEvent) Repeat() (repeat bool) {
	macro.Rewrite("$_.repeat")
	return repeat
}

func (*KeyboardEvent) ShiftKey() (shiftKey bool) {
	macro.Rewrite("$_.shiftKey")
	return shiftKey
}

func (*KeyboardEvent) Which() (which int8) {
	macro.Rewrite("$_.which")
	return which
}

func (*KeyboardEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

func (*KeyboardEvent) View() (view *window.Window) {
	macro.Rewrite("$_.view")
	return view
}

func (*KeyboardEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*KeyboardEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*KeyboardEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*KeyboardEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*KeyboardEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*KeyboardEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*KeyboardEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*KeyboardEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*KeyboardEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*KeyboardEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*KeyboardEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*KeyboardEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*KeyboardEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*KeyboardEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

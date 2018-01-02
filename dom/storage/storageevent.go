package storage

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*StorageEvent)(nil)

type StorageEvent struct {
}

func (*StorageEvent) InitStorageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, keyArg string, oldValueArg interface{}, newValueArg interface{}, urlArg string, storageAreaArg *Storage) {
	macro.Rewrite("$_.initStorageEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, keyArg, oldValueArg, newValueArg, urlArg, storageAreaArg)
}

func (*StorageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*StorageEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*StorageEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*StorageEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*StorageEvent) Key() (key string) {
	macro.Rewrite("$_.key")
	return key
}

func (*StorageEvent) NewValue() (newValue interface{}) {
	macro.Rewrite("$_.newValue")
	return newValue
}

func (*StorageEvent) OldValue() (oldValue interface{}) {
	macro.Rewrite("$_.oldValue")
	return oldValue
}

func (*StorageEvent) StorageArea() (storageArea *Storage) {
	macro.Rewrite("$_.storageArea")
	return storageArea
}

func (*StorageEvent) URL() (url string) {
	macro.Rewrite("$_.url")
	return url
}

func (*StorageEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*StorageEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*StorageEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*StorageEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*StorageEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*StorageEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*StorageEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*StorageEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*StorageEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*StorageEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*StorageEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*StorageEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*StorageEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*StorageEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

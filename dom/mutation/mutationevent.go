package mutation

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*MutationEvent)(nil)

type MutationEvent struct {
}

func (*MutationEvent) InitMutationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, relatedNodeArg dom.Node, prevValueArg string, newValueArg string, attrNameArg string, attrChangeArg uint8) {
	macro.Rewrite("$_.initMutationEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, relatedNodeArg, prevValueArg, newValueArg, attrNameArg, attrChangeArg)
}

func (*MutationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*MutationEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*MutationEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*MutationEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*MutationEvent) AttrChange() (attrChange uint8) {
	macro.Rewrite("$_.attrChange")
	return attrChange
}

func (*MutationEvent) AttrName() (attrName string) {
	macro.Rewrite("$_.attrName")
	return attrName
}

func (*MutationEvent) NewValue() (newValue string) {
	macro.Rewrite("$_.newValue")
	return newValue
}

func (*MutationEvent) PrevValue() (prevValue string) {
	macro.Rewrite("$_.prevValue")
	return prevValue
}

func (*MutationEvent) RelatedNode() (relatedNode dom.Node) {
	macro.Rewrite("$_.relatedNode")
	return relatedNode
}

func (*MutationEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*MutationEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*MutationEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*MutationEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*MutationEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*MutationEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*MutationEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*MutationEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*MutationEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*MutationEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*MutationEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*MutationEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*MutationEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*MutationEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

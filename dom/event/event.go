package event

import "github.com/matthewmueller/joy/dom/element"

type Event interface {
	InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool)

	PreventDefault()

	StopImmediatePropagation()

	StopPropagation()

	Bubbles() (bubbles bool)

	Cancelable() (cancelable bool)

	CancelBubble() (cancelBubble bool)

	SetCancelBubble(cancelBubble bool)

	CurrentTarget() (currentTarget EventTarget)

	DefaultPrevented() (defaultPrevented bool)

	EventPhase() (eventPhase uint8)

	IsTrusted() (isTrusted bool)

	ReturnValue() (returnValue bool)

	SetReturnValue(returnValue bool)

	SrcElement() (srcElement element.Element)

	Target() (target EventTarget)

	TimeStamp() (timeStamp uint64)

	Type() (kind string)
}

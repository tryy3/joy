package mediastreams

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ event.Event = (*MediaStreamTrackEvent)(nil)

func NewMediaStreamTrackEvent(typearg string, eventinitdict *MediaStreamTrackEventInit) *MediaStreamTrackEvent {
	macro.Rewrite("new MediaStreamTrackEvent($1, $2)", typearg, eventinitdict)
	return &MediaStreamTrackEvent{}
}

type MediaStreamTrackEvent struct {
}

func (*MediaStreamTrackEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*MediaStreamTrackEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

func (*MediaStreamTrackEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

func (*MediaStreamTrackEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

func (*MediaStreamTrackEvent) Track() (track *MediaStreamTrack) {
	macro.Rewrite("$_.track")
	return track
}

func (*MediaStreamTrackEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

func (*MediaStreamTrackEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

func (*MediaStreamTrackEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

func (*MediaStreamTrackEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

func (*MediaStreamTrackEvent) CurrentTarget() (currentTarget event.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

func (*MediaStreamTrackEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

func (*MediaStreamTrackEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

func (*MediaStreamTrackEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

func (*MediaStreamTrackEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

func (*MediaStreamTrackEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

func (*MediaStreamTrackEvent) SrcElement() (srcElement element.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

func (*MediaStreamTrackEvent) Target() (target event.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

func (*MediaStreamTrackEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

func (*MediaStreamTrackEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

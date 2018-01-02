package mediastreams

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/webrtc"
)

var _ event.EventTarget = (*MediaDevices)(nil)

type MediaDevices struct {
}

func (*MediaDevices) EnumerateDevices() (m []*webrtc.MediaDeviceInfo) {
	macro.Rewrite("await $_.enumerateDevices()")
	return m
}

func (*MediaDevices) GetSupportedConstraints() (m *MediaTrackSupportedConstraints) {
	macro.Rewrite("$_.getSupportedConstraints()")
	return m
}

func (*MediaDevices) GetUserMedia(constraints *MediaStreamConstraints) (m *MediaStream) {
	macro.Rewrite("await $_.getUserMedia($1)", constraints)
	return m
}

func (*MediaDevices) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaDevices) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MediaDevices) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaDevices) Ondevicechange() (ondevicechange func(event.Event)) {
	macro.Rewrite("$_.ondevicechange")
	return ondevicechange
}

func (*MediaDevices) SetOndevicechange(ondevicechange func(event.Event)) {
	macro.Rewrite("$_.ondevicechange = $1", ondevicechange)
}

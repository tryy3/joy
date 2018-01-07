package window

import (
	"github.com/matthewmueller/joy/dom/mediastreams"
	"github.com/matthewmueller/joy/dom/webrtc"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*MediaDevices)(nil)

// MediaDevices struct
// js:"MediaDevices,omit"
type MediaDevices struct {
}

// EnumerateDevices fn
// js:"enumerateDevices"
func (*MediaDevices) EnumerateDevices() (w []*webrtc.MediaDeviceInfo) {
	macro.Rewrite("await $_.enumerateDevices()")
	return w
}

// GetSupportedConstraints fn
// js:"getSupportedConstraints"
func (*MediaDevices) GetSupportedConstraints() (m *mediastreams.MediaTrackSupportedConstraints) {
	macro.Rewrite("$_.getSupportedConstraints()")
	return m
}

// GetUserMedia fn
// js:"getUserMedia"
func (*MediaDevices) GetUserMedia(constraints *mediastreams.MediaStreamConstraints) (m *MediaStream) {
	macro.Rewrite("await $_.getUserMedia($1)", constraints)
	return m
}

// AddEventListener fn
// js:"addEventListener"
func (*MediaDevices) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MediaDevices) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MediaDevices) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Ondevicechange prop
// js:"ondevicechange"
func (*MediaDevices) Ondevicechange() (ondevicechange func(Event)) {
	macro.Rewrite("$_.ondevicechange")
	return ondevicechange
}

// SetOndevicechange prop
// js:"ondevicechange"
func (*MediaDevices) SetOndevicechange(ondevicechange func(Event)) {
	macro.Rewrite("$_.ondevicechange = $1", ondevicechange)
}

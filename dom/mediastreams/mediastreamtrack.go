package mediastreams

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*MediaStreamTrack)(nil)

type MediaStreamTrack struct {
}

func (*MediaStreamTrack) ApplyConstraints(constraints *MediaTrackConstraints) {
	macro.Rewrite("await $_.applyConstraints($1)", constraints)
}

func (*MediaStreamTrack) Clone() (m *MediaStreamTrack) {
	macro.Rewrite("$_.clone()")
	return m
}

func (*MediaStreamTrack) GetCapabilities() (m *MediaTrackCapabilities) {
	macro.Rewrite("$_.getCapabilities()")
	return m
}

func (*MediaStreamTrack) GetConstraints() (m *MediaTrackConstraints) {
	macro.Rewrite("$_.getConstraints()")
	return m
}

func (*MediaStreamTrack) GetSettings() (m *MediaTrackSettings) {
	macro.Rewrite("$_.getSettings()")
	return m
}

func (*MediaStreamTrack) Stop() {
	macro.Rewrite("$_.stop()")
}

func (*MediaStreamTrack) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaStreamTrack) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MediaStreamTrack) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaStreamTrack) Enabled() (enabled bool) {
	macro.Rewrite("$_.enabled")
	return enabled
}

func (*MediaStreamTrack) SetEnabled(enabled bool) {
	macro.Rewrite("$_.enabled = $1", enabled)
}

func (*MediaStreamTrack) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*MediaStreamTrack) Kind() (kind string) {
	macro.Rewrite("$_.kind")
	return kind
}

func (*MediaStreamTrack) Label() (label string) {
	macro.Rewrite("$_.label")
	return label
}

func (*MediaStreamTrack) Muted() (muted bool) {
	macro.Rewrite("$_.muted")
	return muted
}

func (*MediaStreamTrack) Onended() (onended func(*MediaStreamErrorEvent)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*MediaStreamTrack) SetOnended(onended func(*MediaStreamErrorEvent)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*MediaStreamTrack) Onmute() (onmute func(event.Event)) {
	macro.Rewrite("$_.onmute")
	return onmute
}

func (*MediaStreamTrack) SetOnmute(onmute func(event.Event)) {
	macro.Rewrite("$_.onmute = $1", onmute)
}

func (*MediaStreamTrack) Onoverconstrained() (onoverconstrained func(*MediaStreamErrorEvent)) {
	macro.Rewrite("$_.onoverconstrained")
	return onoverconstrained
}

func (*MediaStreamTrack) SetOnoverconstrained(onoverconstrained func(*MediaStreamErrorEvent)) {
	macro.Rewrite("$_.onoverconstrained = $1", onoverconstrained)
}

func (*MediaStreamTrack) Onunmute() (onunmute func(event.Event)) {
	macro.Rewrite("$_.onunmute")
	return onunmute
}

func (*MediaStreamTrack) SetOnunmute(onunmute func(event.Event)) {
	macro.Rewrite("$_.onunmute = $1", onunmute)
}

func (*MediaStreamTrack) Readonly() (readonly bool) {
	macro.Rewrite("$_.readonly")
	return readonly
}

func (*MediaStreamTrack) ReadyState() (readyState *MediaStreamTrackState) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*MediaStreamTrack) Remote() (remote bool) {
	macro.Rewrite("$_.remote")
	return remote
}

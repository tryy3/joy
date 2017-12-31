package ms


import (
	"github.com/matthewmueller/joy/dom/domerror"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*MSAppAsyncOperation)(nil)

// MSAppAsyncOperation struct
// js:"MSAppAsyncOperation,omit"
type MSAppAsyncOperation struct {
}

// Start fn
// js:"start"
func (*MSAppAsyncOperation) Start() {
	macro.Rewrite("$_.start()")
}

// AddEventListener fn
// js:"addEventListener"
func (*MSAppAsyncOperation) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MSAppAsyncOperation) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MSAppAsyncOperation) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Error prop
// js:"error"
func (*MSAppAsyncOperation) Error() (err *domerror.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

// Oncomplete prop
// js:"oncomplete"
func (*MSAppAsyncOperation) Oncomplete() (oncomplete func(window.Event)) {
	macro.Rewrite("$_.oncomplete")
	return oncomplete
}

// SetOncomplete prop
// js:"oncomplete"
func (*MSAppAsyncOperation) SetOncomplete(oncomplete func(window.Event)) {
	macro.Rewrite("$_.oncomplete = $1", oncomplete)
}

// Onerror prop
// js:"onerror"
func (*MSAppAsyncOperation) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*MSAppAsyncOperation) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// ReadyState prop
// js:"readyState"
func (*MSAppAsyncOperation) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

// Result prop
// js:"result"
func (*MSAppAsyncOperation) Result() (result interface{}) {
	macro.Rewrite("$_.result")
	return result
}

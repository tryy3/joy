package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
)

var _ event.EventTarget = (*MSAppAsyncOperation)(nil)

type MSAppAsyncOperation struct {
}

func (*MSAppAsyncOperation) Start() {
	macro.Rewrite("$_.start()")
}

func (*MSAppAsyncOperation) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MSAppAsyncOperation) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MSAppAsyncOperation) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MSAppAsyncOperation) Error() (err *dom.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

func (*MSAppAsyncOperation) Oncomplete() (oncomplete func(event.Event)) {
	macro.Rewrite("$_.oncomplete")
	return oncomplete
}

func (*MSAppAsyncOperation) SetOncomplete(oncomplete func(event.Event)) {
	macro.Rewrite("$_.oncomplete = $1", oncomplete)
}

func (*MSAppAsyncOperation) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*MSAppAsyncOperation) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*MSAppAsyncOperation) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*MSAppAsyncOperation) Result() (result interface{}) {
	macro.Rewrite("$_.result")
	return result
}

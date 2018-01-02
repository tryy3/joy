package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
)

var _ event.EventTarget = (*MSWebViewAsyncOperation)(nil)

type MSWebViewAsyncOperation struct {
}

func (*MSWebViewAsyncOperation) Start() {
	macro.Rewrite("$_.start()")
}

func (*MSWebViewAsyncOperation) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MSWebViewAsyncOperation) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MSWebViewAsyncOperation) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MSWebViewAsyncOperation) Error() (err *dom.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

func (*MSWebViewAsyncOperation) Oncomplete() (oncomplete func(event.Event)) {
	macro.Rewrite("$_.oncomplete")
	return oncomplete
}

func (*MSWebViewAsyncOperation) SetOncomplete(oncomplete func(event.Event)) {
	macro.Rewrite("$_.oncomplete = $1", oncomplete)
}

func (*MSWebViewAsyncOperation) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*MSWebViewAsyncOperation) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*MSWebViewAsyncOperation) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*MSWebViewAsyncOperation) Result() (result interface{}) {
	macro.Rewrite("$_.result")
	return result
}

func (*MSWebViewAsyncOperation) Target() (target *MSHTMLWebViewElement) {
	macro.Rewrite("$_.target")
	return target
}

func (*MSWebViewAsyncOperation) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}

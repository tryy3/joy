package worker

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/utils"
)

var _ event.EventTarget = (*Worker)(nil)

func New(stringurl string) *Worker {
	macro.Rewrite("new Worker($1)", stringurl)
	return &Worker{}
}

type Worker struct {
}

func (*Worker) PostMessage(message interface{}, transfer *[]interface{}) {
	macro.Rewrite("$_.postMessage($1, $2)", message, transfer)
}

func (*Worker) Terminate() {
	macro.Rewrite("$_.terminate()")
}

func (*Worker) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Worker) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*Worker) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Worker) Onmessage() (onmessage func(*utils.MessageEvent)) {
	macro.Rewrite("$_.onmessage")
	return onmessage
}

func (*Worker) SetOnmessage(onmessage func(*utils.MessageEvent)) {
	macro.Rewrite("$_.onmessage = $1", onmessage)
}

func (*Worker) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*Worker) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

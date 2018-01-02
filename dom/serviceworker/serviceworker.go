package serviceworker

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*ServiceWorker)(nil)

type ServiceWorker struct {
}

func (*ServiceWorker) PostMessage(message interface{}, transfer *[]interface{}) {
	macro.Rewrite("$_.postMessage($1, $2)", message, transfer)
}

func (*ServiceWorker) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ServiceWorker) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*ServiceWorker) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ServiceWorker) Onstatechange() (onstatechange func(event.Event)) {
	macro.Rewrite("$_.onstatechange")
	return onstatechange
}

func (*ServiceWorker) SetOnstatechange(onstatechange func(event.Event)) {
	macro.Rewrite("$_.onstatechange = $1", onstatechange)
}

func (*ServiceWorker) ScriptURL() (scriptURL string) {
	macro.Rewrite("$_.scriptURL")
	return scriptURL
}

func (*ServiceWorker) State() (state *ServiceWorkerState) {
	macro.Rewrite("$_.state")
	return state
}

func (*ServiceWorker) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*ServiceWorker) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

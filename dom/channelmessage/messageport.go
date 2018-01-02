package channelmessage

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*MessagePort)(nil)

type MessagePort struct {
}

func (*MessagePort) Close() {
	macro.Rewrite("$_.close()")
}

func (*MessagePort) PostMessage(message *interface{}, transfer *[]interface{}) {
	macro.Rewrite("$_.postMessage($1, $2)", message, transfer)
}

func (*MessagePort) Start() {
	macro.Rewrite("$_.start()")
}

func (*MessagePort) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MessagePort) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MessagePort) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MessagePort) Onmessage() (onmessage func(*utils.MessageEvent)) {
	macro.Rewrite("$_.onmessage")
	return onmessage
}

func (*MessagePort) SetOnmessage(onmessage func(*utils.MessageEvent)) {
	macro.Rewrite("$_.onmessage = $1", onmessage)
}

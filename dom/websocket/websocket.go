package websocket

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/utils"
)

var _ event.EventTarget = (*WebSocket)(nil)

func New(url string, protocols *interface{}) *WebSocket {
	macro.Rewrite("new WebSocket($1, $2)", url, protocols)
	return &WebSocket{}
}

type WebSocket struct {
}

func (*WebSocket) Close(code *uint8, reason *string) {
	macro.Rewrite("$_.close($1, $2)", code, reason)
}

func (*WebSocket) Send(data interface{}) {
	macro.Rewrite("$_.send($1)", data)
}

func (*WebSocket) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*WebSocket) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*WebSocket) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*WebSocket) BinaryType() (binaryType string) {
	macro.Rewrite("$_.binaryType")
	return binaryType
}

func (*WebSocket) SetBinaryType(binaryType string) {
	macro.Rewrite("$_.binaryType = $1", binaryType)
}

func (*WebSocket) BufferedAmount() (bufferedAmount uint) {
	macro.Rewrite("$_.bufferedAmount")
	return bufferedAmount
}

func (*WebSocket) Extensions() (extensions string) {
	macro.Rewrite("$_.extensions")
	return extensions
}

func (*WebSocket) Onclose() (onclose func(*CloseEvent)) {
	macro.Rewrite("$_.onclose")
	return onclose
}

func (*WebSocket) SetOnclose(onclose func(*CloseEvent)) {
	macro.Rewrite("$_.onclose = $1", onclose)
}

func (*WebSocket) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*WebSocket) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*WebSocket) Onmessage() (onmessage func(*utils.MessageEvent)) {
	macro.Rewrite("$_.onmessage")
	return onmessage
}

func (*WebSocket) SetOnmessage(onmessage func(*utils.MessageEvent)) {
	macro.Rewrite("$_.onmessage = $1", onmessage)
}

func (*WebSocket) Onopen() (onopen func(event.Event)) {
	macro.Rewrite("$_.onopen")
	return onopen
}

func (*WebSocket) SetOnopen(onopen func(event.Event)) {
	macro.Rewrite("$_.onopen = $1", onopen)
}

func (*WebSocket) Protocol() (protocol string) {
	macro.Rewrite("$_.protocol")
	return protocol
}

func (*WebSocket) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*WebSocket) URL() (url string) {
	macro.Rewrite("$_.url")
	return url
}

package indexdb

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
)

var _ IDBRequest = (*IDBOpenDBRequest)(nil)
var _ event.EventTarget = (*IDBOpenDBRequest)(nil)

type IDBOpenDBRequest struct {
}

func (*IDBOpenDBRequest) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*IDBOpenDBRequest) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*IDBOpenDBRequest) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*IDBOpenDBRequest) Onblocked() (onblocked func(event.Event)) {
	macro.Rewrite("$_.onblocked")
	return onblocked
}

func (*IDBOpenDBRequest) SetOnblocked(onblocked func(event.Event)) {
	macro.Rewrite("$_.onblocked = $1", onblocked)
}

func (*IDBOpenDBRequest) Onupgradeneeded() (onupgradeneeded func(*IDBVersionChangeEvent)) {
	macro.Rewrite("$_.onupgradeneeded")
	return onupgradeneeded
}

func (*IDBOpenDBRequest) SetOnupgradeneeded(onupgradeneeded func(*IDBVersionChangeEvent)) {
	macro.Rewrite("$_.onupgradeneeded = $1", onupgradeneeded)
}

func (*IDBOpenDBRequest) Error() (err *dom.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

func (*IDBOpenDBRequest) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*IDBOpenDBRequest) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*IDBOpenDBRequest) Onsuccess() (onsuccess func(event.Event)) {
	macro.Rewrite("$_.onsuccess")
	return onsuccess
}

func (*IDBOpenDBRequest) SetOnsuccess(onsuccess func(event.Event)) {
	macro.Rewrite("$_.onsuccess = $1", onsuccess)
}

func (*IDBOpenDBRequest) ReadyState() (readyState *IDBRequestReadyState) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*IDBOpenDBRequest) Result() (result interface{}) {
	macro.Rewrite("$_.result")
	return result
}

func (*IDBOpenDBRequest) Source() (source interface{}) {
	macro.Rewrite("$_.source")
	return source
}

func (*IDBOpenDBRequest) Transaction() (transaction *IDBTransaction) {
	macro.Rewrite("$_.transaction")
	return transaction
}

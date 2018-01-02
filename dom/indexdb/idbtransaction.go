package indexdb

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
)

var _ event.EventTarget = (*IDBTransaction)(nil)

type IDBTransaction struct {
}

func (*IDBTransaction) Abort() {
	macro.Rewrite("$_.abort()")
}

func (*IDBTransaction) ObjectStore(name string) (i *IDBObjectStore) {
	macro.Rewrite("$_.objectStore($1)", name)
	return i
}

func (*IDBTransaction) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*IDBTransaction) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*IDBTransaction) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*IDBTransaction) Db() (db *IDBDatabase) {
	macro.Rewrite("$_.db")
	return db
}

func (*IDBTransaction) Error() (err *dom.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

func (*IDBTransaction) Mode() (mode *IDBTransactionMode) {
	macro.Rewrite("$_.mode")
	return mode
}

func (*IDBTransaction) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*IDBTransaction) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*IDBTransaction) Oncomplete() (oncomplete func(event.Event)) {
	macro.Rewrite("$_.oncomplete")
	return oncomplete
}

func (*IDBTransaction) SetOncomplete(oncomplete func(event.Event)) {
	macro.Rewrite("$_.oncomplete = $1", oncomplete)
}

func (*IDBTransaction) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*IDBTransaction) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

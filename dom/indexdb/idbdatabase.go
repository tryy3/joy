package indexdb

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
)

var _ event.EventTarget = (*IDBDatabase)(nil)

type IDBDatabase struct {
}

func (*IDBDatabase) Close() {
	macro.Rewrite("$_.close()")
}

func (*IDBDatabase) CreateObjectStore(name string, optionalParameters *IDBObjectStoreParameters) (i *IDBObjectStore) {
	macro.Rewrite("$_.createObjectStore($1, $2)", name, optionalParameters)
	return i
}

func (*IDBDatabase) DeleteObjectStore(name string) {
	macro.Rewrite("$_.deleteObjectStore($1)", name)
}

func (*IDBDatabase) Transaction(storeNames interface{}, mode *IDBTransactionMode) (i *IDBTransaction) {
	macro.Rewrite("$_.transaction($1, $2)", storeNames, mode)
	return i
}

func (*IDBDatabase) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*IDBDatabase) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*IDBDatabase) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*IDBDatabase) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*IDBDatabase) ObjectStoreNames() (objectStoreNames *dom.DOMStringList) {
	macro.Rewrite("$_.objectStoreNames")
	return objectStoreNames
}

func (*IDBDatabase) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*IDBDatabase) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*IDBDatabase) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*IDBDatabase) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*IDBDatabase) Version() (version uint64) {
	macro.Rewrite("$_.version")
	return version
}

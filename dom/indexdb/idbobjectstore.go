package indexdb

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
)

type IDBObjectStore struct {
}

func (*IDBObjectStore) Add(value interface{}, key *interface{}) (i IDBRequest) {
	macro.Rewrite("$_.add($1, $2)", value, key)
	return i
}

func (*IDBObjectStore) Clear() (i IDBRequest) {
	macro.Rewrite("$_.clear()")
	return i
}

func (*IDBObjectStore) Count(key *interface{}) (i IDBRequest) {
	macro.Rewrite("$_.count($1)", key)
	return i
}

func (*IDBObjectStore) CreateIndex(name string, keyPath string, optionalParameters *IDBIndexParameters) (i *IDBIndex) {
	macro.Rewrite("$_.createIndex($1, $2, $3)", name, keyPath, optionalParameters)
	return i
}

func (*IDBObjectStore) Delete(key interface{}) (i IDBRequest) {
	macro.Rewrite("$_.delete($1)", key)
	return i
}

func (*IDBObjectStore) DeleteIndex(indexName string) {
	macro.Rewrite("$_.deleteIndex($1)", indexName)
}

func (*IDBObjectStore) Get(key interface{}) (i IDBRequest) {
	macro.Rewrite("$_.get($1)", key)
	return i
}

func (*IDBObjectStore) Index(name string) (i *IDBIndex) {
	macro.Rewrite("$_.index($1)", name)
	return i
}

func (*IDBObjectStore) OpenCursor(rng *interface{}, direction *IDBCursorDirection) (i IDBRequest) {
	macro.Rewrite("$_.openCursor($1, $2)", rng, direction)
	return i
}

func (*IDBObjectStore) Put(value interface{}, key *interface{}) (i IDBRequest) {
	macro.Rewrite("$_.put($1, $2)", value, key)
	return i
}

func (*IDBObjectStore) IndexNames() (indexNames *dom.DOMStringList) {
	macro.Rewrite("$_.indexNames")
	return indexNames
}

func (*IDBObjectStore) KeyPath() (keyPath string) {
	macro.Rewrite("$_.keyPath")
	return keyPath
}

func (*IDBObjectStore) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*IDBObjectStore) Transaction() (transaction *IDBTransaction) {
	macro.Rewrite("$_.transaction")
	return transaction
}

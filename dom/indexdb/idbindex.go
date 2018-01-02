package indexdb

import (
	"github.com/matthewmueller/joy/macro"
)

type IDBIndex struct {
}

func (*IDBIndex) Count(key *interface{}) (i IDBRequest) {
	macro.Rewrite("$_.count($1)", key)
	return i
}

func (*IDBIndex) Get(key interface{}) (i IDBRequest) {
	macro.Rewrite("$_.get($1)", key)
	return i
}

func (*IDBIndex) GetKey(key interface{}) (i IDBRequest) {
	macro.Rewrite("$_.getKey($1)", key)
	return i
}

func (*IDBIndex) OpenCursor(rng *interface{}, direction *IDBCursorDirection) (i IDBRequest) {
	macro.Rewrite("$_.openCursor($1, $2)", rng, direction)
	return i
}

func (*IDBIndex) OpenKeyCursor(rng *interface{}, direction *IDBCursorDirection) (i IDBRequest) {
	macro.Rewrite("$_.openKeyCursor($1, $2)", rng, direction)
	return i
}

func (*IDBIndex) KeyPath() (keyPath string) {
	macro.Rewrite("$_.keyPath")
	return keyPath
}

func (*IDBIndex) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*IDBIndex) ObjectStore() (objectStore *IDBObjectStore) {
	macro.Rewrite("$_.objectStore")
	return objectStore
}

func (*IDBIndex) Unique() (unique bool) {
	macro.Rewrite("$_.unique")
	return unique
}

package window

import (
	"github.com/matthewmueller/joy/dom/indexdb"
	"github.com/matthewmueller/joy/macro"
)

// IDBIndex struct
// js:"IDBIndex,omit"
type IDBIndex struct {
}

// Count fn
// js:"count"
func (*IDBIndex) Count(key *interface{}) (i IDBRequest) {
	macro.Rewrite("$_.count($1)", key)
	return i
}

// Get fn
// js:"get"
func (*IDBIndex) Get(key interface{}) (i IDBRequest) {
	macro.Rewrite("$_.get($1)", key)
	return i
}

// GetKey fn
// js:"getKey"
func (*IDBIndex) GetKey(key interface{}) (i IDBRequest) {
	macro.Rewrite("$_.getKey($1)", key)
	return i
}

// OpenCursor fn
// js:"openCursor"
func (*IDBIndex) OpenCursor(rng *interface{}, direction *indexdb.IDBCursorDirection) (i IDBRequest) {
	macro.Rewrite("$_.openCursor($1, $2)", rng, direction)
	return i
}

// OpenKeyCursor fn
// js:"openKeyCursor"
func (*IDBIndex) OpenKeyCursor(rng *interface{}, direction *indexdb.IDBCursorDirection) (i IDBRequest) {
	macro.Rewrite("$_.openKeyCursor($1, $2)", rng, direction)
	return i
}

// KeyPath prop
// js:"keyPath"
func (*IDBIndex) KeyPath() (keyPath string) {
	macro.Rewrite("$_.keyPath")
	return keyPath
}

// Name prop
// js:"name"
func (*IDBIndex) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// ObjectStore prop
// js:"objectStore"
func (*IDBIndex) ObjectStore() (objectStore *IDBObjectStore) {
	macro.Rewrite("$_.objectStore")
	return objectStore
}

// Unique prop
// js:"unique"
func (*IDBIndex) Unique() (unique bool) {
	macro.Rewrite("$_.unique")
	return unique
}

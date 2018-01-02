package indexdb

import (
	"github.com/matthewmueller/joy/macro"
)

var _ IDBCursor = (*IDBCursorWithValue)(nil)

type IDBCursorWithValue struct {
}

func (*IDBCursorWithValue) Advance(count uint) {
	macro.Rewrite("$_.advance($1)", count)
}

func (*IDBCursorWithValue) Continue(key *interface{}) {
	macro.Rewrite("$_.continue($1)", key)
}

func (*IDBCursorWithValue) Delete() (w IDBRequest) {
	macro.Rewrite("$_.delete()")
	return w
}

func (*IDBCursorWithValue) Update(value interface{}) (w IDBRequest) {
	macro.Rewrite("$_.update($1)", value)
	return w
}

func (*IDBCursorWithValue) Value() (value interface{}) {
	macro.Rewrite("$_.value")
	return value
}

func (*IDBCursorWithValue) Direction() (direction *IDBCursorDirection) {
	macro.Rewrite("$_.direction")
	return direction
}

func (*IDBCursorWithValue) Key() (key interface{}) {
	macro.Rewrite("$_.key")
	return key
}

func (*IDBCursorWithValue) PrimaryKey() (primaryKey interface{}) {
	macro.Rewrite("$_.primaryKey")
	return primaryKey
}

func (*IDBCursorWithValue) Source() (source interface{}) {
	macro.Rewrite("$_.source")
	return source
}

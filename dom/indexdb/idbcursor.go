package indexdb

type IDBCursor interface {
	Advance(count uint)

	Continue(key *interface{})

	Delete() (w IDBRequest)

	Update(value interface{}) (w IDBRequest)

	Direction() (direction *IDBCursorDirection)

	Key() (key interface{})

	PrimaryKey() (primaryKey interface{})

	Source() (source interface{})
}

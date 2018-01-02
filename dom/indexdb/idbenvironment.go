package indexdb

type IDBEnvironment interface {
	IndexedDB() (indexedDB *IDBFactory)
}

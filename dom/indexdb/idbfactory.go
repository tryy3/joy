package indexdb

import "github.com/matthewmueller/joy/macro"

type IDBFactory struct {
}

func (*IDBFactory) Cmp(first interface{}, second interface{}) (i int8) {
	macro.Rewrite("$_.cmp($1, $2)", first, second)
	return i
}

func (*IDBFactory) DeleteDatabase(name string) (i *IDBOpenDBRequest) {
	macro.Rewrite("$_.deleteDatabase($1)", name)
	return i
}

func (*IDBFactory) Open(name string, version *uint64) (i *IDBOpenDBRequest) {
	macro.Rewrite("$_.open($1, $2)", name, version)
	return i
}

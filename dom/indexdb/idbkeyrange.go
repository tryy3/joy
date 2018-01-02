package indexdb

import "github.com/matthewmueller/joy/macro"

type IDBKeyRange struct {
}

func (*IDBKeyRange) Bound(lower interface{}, upper interface{}, lowerOpen *bool, upperOpen *bool) (i *IDBKeyRange) {
	macro.Rewrite("$_.bound($1, $2, $3, $4)", lower, upper, lowerOpen, upperOpen)
	return i
}

func (*IDBKeyRange) LowerBound(lower interface{}, open *bool) (i *IDBKeyRange) {
	macro.Rewrite("$_.lowerBound($1, $2)", lower, open)
	return i
}

func (*IDBKeyRange) Only(value interface{}) (i *IDBKeyRange) {
	macro.Rewrite("$_.only($1)", value)
	return i
}

func (*IDBKeyRange) UpperBound(upper interface{}, open *bool) (i *IDBKeyRange) {
	macro.Rewrite("$_.upperBound($1, $2)", upper, open)
	return i
}

func (*IDBKeyRange) Lower() (lower interface{}) {
	macro.Rewrite("$_.lower")
	return lower
}

func (*IDBKeyRange) LowerOpen() (lowerOpen bool) {
	macro.Rewrite("$_.lowerOpen")
	return lowerOpen
}

func (*IDBKeyRange) Upper() (upper interface{}) {
	macro.Rewrite("$_.upper")
	return upper
}

func (*IDBKeyRange) UpperOpen() (upperOpen bool) {
	macro.Rewrite("$_.upperOpen")
	return upperOpen
}

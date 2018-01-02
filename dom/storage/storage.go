package storage

import "github.com/matthewmueller/joy/macro"

type Storage struct {
}

func (*Storage) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*Storage) GetItem(key string) (i interface{}) {
	macro.Rewrite("$_.getItem($1)", key)
	return i
}

func (*Storage) Key(index uint) (s string) {
	macro.Rewrite("$_.key($1)", index)
	return s
}

func (*Storage) RemoveItem(key string) {
	macro.Rewrite("$_.removeItem($1)", key)
}

func (*Storage) SetItem(key string, data string) {
	macro.Rewrite("$_.setItem($1, $2)", key, data)
}

func (*Storage) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

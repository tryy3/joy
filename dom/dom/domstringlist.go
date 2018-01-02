package dom

import "github.com/matthewmueller/joy/macro"

type DOMStringList struct {
}

func (*DOMStringList) Contains(str string) (b bool) {
	macro.Rewrite("$_.contains($1)", str)
	return b
}

func (*DOMStringList) Item(index uint) (s string) {
	macro.Rewrite("$_.item($1)", index)
	return s
}

func (*DOMStringList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

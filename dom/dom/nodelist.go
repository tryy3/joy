package dom

import "github.com/matthewmueller/joy/macro"

type NodeList struct {
}

func (*NodeList) Item(index uint) (n Node) {
	macro.Rewrite("$_.item($1)", index)
	return n
}

func (*NodeList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

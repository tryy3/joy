package touch

import "github.com/matthewmueller/joy/macro"

type TouchList struct {
}

func (*TouchList) Item(index uint) (t *Touch) {
	macro.Rewrite("$_.item($1)", index)
	return t
}

func (*TouchList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

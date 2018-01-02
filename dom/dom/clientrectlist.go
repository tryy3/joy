package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type ClientRectList struct {
}

func (*ClientRectList) Item(index uint) (c *ClientRect) {
	macro.Rewrite("$_.item($1)", index)
	return c
}

func (*ClientRectList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

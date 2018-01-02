package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
)

type MSRangeCollection struct {
}

func (*MSRangeCollection) Item(index uint) (r *dom.Range) {
	macro.Rewrite("$_.item($1)", index)
	return r
}

func (*MSRangeCollection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

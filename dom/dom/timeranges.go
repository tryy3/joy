package dom

import "github.com/matthewmueller/joy/macro"

type TimeRanges struct {
}

func (*TimeRanges) End(index uint) (f float32) {
	macro.Rewrite("$_.end($1)", index)
	return f
}

func (*TimeRanges) Start(index uint) (f float32) {
	macro.Rewrite("$_.start($1)", index)
	return f
}

func (*TimeRanges) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

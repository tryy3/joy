package file

import (
	"github.com/matthewmueller/joy/macro"
)

type FileList struct {
}

func (*FileList) Item(index uint) (f *File) {
	macro.Rewrite("$_.item($1)", index)
	return f
}

func (*FileList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

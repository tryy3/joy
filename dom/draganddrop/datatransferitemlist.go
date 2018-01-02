package draganddrop

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/file"
)

type DataTransferItemList struct {
}

func (*DataTransferItemList) Add(data *file.File) (d *DataTransferItem) {
	macro.Rewrite("$_.add($1)", data)
	return d
}

func (*DataTransferItemList) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*DataTransferItemList) Item(index uint) (f *file.File) {
	macro.Rewrite("$_.item($1)", index)
	return f
}

func (*DataTransferItemList) Remove(index uint) {
	macro.Rewrite("$_.remove($1)", index)
}

func (*DataTransferItemList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

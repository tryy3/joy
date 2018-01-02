package draganddrop

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/file"
)

type DataTransferItem struct {
}

func (*DataTransferItem) GetAsFile() (f *file.File) {
	macro.Rewrite("$_.getAsFile()")
	return f
}

func (*DataTransferItem) GetAsString(Callback func(data string)) {
	macro.Rewrite("$_.getAsString($1)", Callback)
}

func (*DataTransferItem) WebkitGetAsEntry() (i interface{}) {
	macro.Rewrite("$_.webkitGetAsEntry()")
	return i
}

func (*DataTransferItem) Kind() (kind string) {
	macro.Rewrite("$_.kind")
	return kind
}

func (*DataTransferItem) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

package draganddrop

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/dom/dom"
)

type DataTransfer struct {
}

func (*DataTransfer) ClearData(format *string) (b bool) {
	macro.Rewrite("$_.clearData($1)", format)
	return b
}

func (*DataTransfer) GetData(format string) (s string) {
	macro.Rewrite("$_.getData($1)", format)
	return s
}

func (*DataTransfer) SetData(format string, data string) (b bool) {
	macro.Rewrite("$_.setData($1, $2)", format, data)
	return b
}

func (*DataTransfer) DropEffect() (dropEffect string) {
	macro.Rewrite("$_.dropEffect")
	return dropEffect
}

func (*DataTransfer) SetDropEffect(dropEffect string) {
	macro.Rewrite("$_.dropEffect = $1", dropEffect)
}

func (*DataTransfer) EffectAllowed() (effectAllowed string) {
	macro.Rewrite("$_.effectAllowed")
	return effectAllowed
}

func (*DataTransfer) SetEffectAllowed(effectAllowed string) {
	macro.Rewrite("$_.effectAllowed = $1", effectAllowed)
}

func (*DataTransfer) Files() (files *file.FileList) {
	macro.Rewrite("$_.files")
	return files
}

func (*DataTransfer) Items() (items *DataTransferItemList) {
	macro.Rewrite("$_.items")
	return items
}

func (*DataTransfer) Types() (types *dom.DOMStringList) {
	macro.Rewrite("$_.types")
	return types
}

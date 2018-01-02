package xhr

import "github.com/matthewmueller/joy/macro"

func New() *FormData {
	macro.Rewrite("new FormData()")
	return &FormData{}
}

type FormData struct {
}

func (*FormData) Append(name interface{}, value interface{}, blobName *string) {
	macro.Rewrite("$_.append($1, $2, $3)", name, value, blobName)
}

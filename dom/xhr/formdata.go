package xhr

import "github.com/matthewmueller/joy/macro"

// NewFormData fn
func NewFormData() *FormData {
	macro.Rewrite("new FormData()")
	return &FormData{}
}

// FormData struct
// js:"FormData,omit"
type FormData struct {
}

// Append fn
// js:"append"
func (*FormData) Append(name interface{}, value interface{}, blobName *string) {
	macro.Rewrite("$_.append($1, $2, $3)", name, value, blobName)
}

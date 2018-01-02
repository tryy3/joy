package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/file"
)

func New() *MSBlobBuilder {
	macro.Rewrite("new MSBlobBuilder()")
	return &MSBlobBuilder{}
}

type MSBlobBuilder struct {
}

func (*MSBlobBuilder) Append(data interface{}, endings *string) {
	macro.Rewrite("$_.append($1, $2)", data, endings)
}

func (*MSBlobBuilder) GetBlob(contentType *string) (b file.Blob) {
	macro.Rewrite("$_.getBlob($1)", contentType)
	return b
}

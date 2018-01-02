package file

import (
	"github.com/matthewmueller/joy/macro"
)

var _ Blob = (*File)(nil)

type File struct {
}

func (*File) MsClose() {
	macro.Rewrite("$_.msClose()")
}

func (*File) MsDetachStream() (i interface{}) {
	macro.Rewrite("$_.msDetachStream()")
	return i
}

func (*File) Slice(start *int64, end *int64, contentType *string) (b Blob) {
	macro.Rewrite("$_.slice($1, $2, $3)", start, end, contentType)
	return b
}

func (*File) LastModifiedDate() (lastModifiedDate interface{}) {
	macro.Rewrite("$_.lastModifiedDate")
	return lastModifiedDate
}

func (*File) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*File) WebkitRelativePath() (webkitRelativePath string) {
	macro.Rewrite("$_.webkitRelativePath")
	return webkitRelativePath
}

func (*File) Size() (size uint64) {
	macro.Rewrite("$_.size")
	return size
}

func (*File) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

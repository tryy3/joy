package streams

import (
	"github.com/matthewmueller/joy/macro"
)

type ReadableStream struct {
}

func (*ReadableStream) Cancel() {
	macro.Rewrite("await $_.cancel()")
}

func (*ReadableStream) GetReader() (r *ReadableStreamReader) {
	macro.Rewrite("$_.getReader()")
	return r
}

func (*ReadableStream) Locked() (locked bool) {
	macro.Rewrite("$_.locked")
	return locked
}

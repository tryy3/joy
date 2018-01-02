package streams

import "github.com/matthewmueller/joy/macro"

type ReadableStreamReader struct {
}

func (*ReadableStreamReader) Cancel() {
	macro.Rewrite("await $_.cancel()")
}

func (*ReadableStreamReader) Read() (i interface{}) {
	macro.Rewrite("await $_.read()")
	return i
}

func (*ReadableStreamReader) ReleaseLock() {
	macro.Rewrite("$_.releaseLock()")
}

package ms

import "github.com/matthewmueller/joy/macro"

type MSStream struct {
}

func (*MSStream) MsClose() {
	macro.Rewrite("$_.msClose()")
}

func (*MSStream) MsDetachStream() (i interface{}) {
	macro.Rewrite("$_.msDetachStream()")
	return i
}

func (*MSStream) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

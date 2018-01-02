package webkit

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
)

type WebKitDirectoryReader struct {
}

func (*WebKitDirectoryReader) ReadEntries(successCallback func(entries []*WebKitEntry), errorCallback *func(err *dom.DOMError)) {
	macro.Rewrite("$_.readEntries($1, $2)", successCallback, errorCallback)
}

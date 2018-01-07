package webkitdirectoryreader

import (
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/macro"
)

// WebKitDirectoryReader struct
// js:"WebKitDirectoryReader,omit"
type WebKitDirectoryReader struct {
}

// ReadEntries fn
// js:"readEntries"
func (*WebKitDirectoryReader) ReadEntries(successCallback func(entries []*WebKitEntry), errorCallback *func(err *dom.DOMError)) {
	macro.Rewrite("$_.readEntries($1, $2)", successCallback, errorCallback)
}

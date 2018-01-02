package webkit

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/dom/dom"
)

var _ WebKitEntry = (*WebKitFileEntry)(nil)

type WebKitFileEntry struct {
}

func (*WebKitFileEntry) File(successCallback func(file *file.File), errorCallback *func(err *dom.DOMError)) {
	macro.Rewrite("$_.file($1, $2)", successCallback, errorCallback)
}

func (*WebKitFileEntry) Filesystem() (filesystem *WebKitFileSystem) {
	macro.Rewrite("$_.filesystem")
	return filesystem
}

func (*WebKitFileEntry) FullPath() (fullPath string) {
	macro.Rewrite("$_.fullPath")
	return fullPath
}

func (*WebKitFileEntry) IsDirectory() (isDirectory bool) {
	macro.Rewrite("$_.isDirectory")
	return isDirectory
}

func (*WebKitFileEntry) IsFile() (isFile bool) {
	macro.Rewrite("$_.isFile")
	return isFile
}

func (*WebKitFileEntry) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

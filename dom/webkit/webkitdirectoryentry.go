package webkit

import (
	"github.com/matthewmueller/joy/macro"
)

var _ WebKitEntry = (*WebKitDirectoryEntry)(nil)

type WebKitDirectoryEntry struct {
}

func (*WebKitDirectoryEntry) CreateReader() (w *WebKitDirectoryReader) {
	macro.Rewrite("$_.createReader()")
	return w
}

func (*WebKitDirectoryEntry) Filesystem() (filesystem *WebKitFileSystem) {
	macro.Rewrite("$_.filesystem")
	return filesystem
}

func (*WebKitDirectoryEntry) FullPath() (fullPath string) {
	macro.Rewrite("$_.fullPath")
	return fullPath
}

func (*WebKitDirectoryEntry) IsDirectory() (isDirectory bool) {
	macro.Rewrite("$_.isDirectory")
	return isDirectory
}

func (*WebKitDirectoryEntry) IsFile() (isFile bool) {
	macro.Rewrite("$_.isFile")
	return isFile
}

func (*WebKitDirectoryEntry) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

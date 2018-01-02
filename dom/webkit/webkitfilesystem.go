package webkit

import "github.com/matthewmueller/joy/macro"

type WebKitFileSystem struct {
}

func (*WebKitFileSystem) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*WebKitFileSystem) Root() (root *WebKitDirectoryEntry) {
	macro.Rewrite("$_.root")
	return root
}

package webkit

type WebKitEntry interface {
	Filesystem() (filesystem *WebKitFileSystem)

	FullPath() (fullPath string)

	IsDirectory() (isDirectory bool)

	IsFile() (isFile bool)

	Name() (name string)
}

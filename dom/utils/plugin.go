package utils

import "github.com/matthewmueller/joy/macro"

type Plugin struct {
}

func (*Plugin) Item(index uint) (m *MimeType) {
	macro.Rewrite("$_.item($1)", index)
	return m
}

func (*Plugin) NamedItem(kind string) (m *MimeType) {
	macro.Rewrite("$_.namedItem($1)", kind)
	return m
}

func (*Plugin) Description() (description string) {
	macro.Rewrite("$_.description")
	return description
}

func (*Plugin) Filename() (filename string) {
	macro.Rewrite("$_.filename")
	return filename
}

func (*Plugin) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*Plugin) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*Plugin) Version() (version string) {
	macro.Rewrite("$_.version")
	return version
}

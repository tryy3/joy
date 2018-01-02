package utils

import (
	"github.com/matthewmueller/joy/macro"
)

type MimeTypeArray struct {
}

func (*MimeTypeArray) Item(index uint) (m *Plugin) {
	macro.Rewrite("$_.item($1)", index)
	return m
}

func (*MimeTypeArray) NamedItem(kind string) (m *Plugin) {
	macro.Rewrite("$_.namedItem($1)", kind)
	return m
}

func (*MimeTypeArray) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

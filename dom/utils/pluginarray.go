package utils

import (
	"github.com/matthewmueller/joy/macro"
)

type PluginArray struct {
}

func (*PluginArray) Item(index uint) (m *Plugin) {
	macro.Rewrite("$_.item($1)", index)
	return m
}

func (*PluginArray) NamedItem(name string) (m *Plugin) {
	macro.Rewrite("$_.namedItem($1)", name)
	return m
}

func (*PluginArray) Refresh(reload *bool) {
	macro.Rewrite("$_.refresh($1)", reload)
}

func (*PluginArray) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

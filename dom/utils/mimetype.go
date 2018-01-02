package utils

import "github.com/matthewmueller/joy/macro"

type MimeType struct {
}

func (*MimeType) Description() (description string) {
	macro.Rewrite("$_.description")
	return description
}

func (*MimeType) EnabledPlugin() (enabledPlugin *Plugin) {
	macro.Rewrite("$_.enabledPlugin")
	return enabledPlugin
}

func (*MimeType) Suffixes() (suffixes string) {
	macro.Rewrite("$_.suffixes")
	return suffixes
}

func (*MimeType) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

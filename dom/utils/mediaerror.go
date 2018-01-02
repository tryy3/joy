package utils

import "github.com/matthewmueller/joy/macro"

type MediaError struct {
}

func (*MediaError) Code() (code int8) {
	macro.Rewrite("$_.code")
	return code
}

func (*MediaError) MsExtendedCode() (msExtendedCode int) {
	macro.Rewrite("$_.msExtendedCode")
	return msExtendedCode
}

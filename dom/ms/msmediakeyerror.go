package ms

import "github.com/matthewmueller/joy/macro"

type MSMediaKeyError struct {
}

func (*MSMediaKeyError) Code() (code uint8) {
	macro.Rewrite("$_.code")
	return code
}

func (*MSMediaKeyError) SystemCode() (systemCode uint) {
	macro.Rewrite("$_.systemCode")
	return systemCode
}

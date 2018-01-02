package geolocation

import "github.com/matthewmueller/joy/macro"

type PositionError struct {
}

func (*PositionError) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

func (*PositionError) Code() (code uint8) {
	macro.Rewrite("$_.code")
	return code
}

func (*PositionError) Message() (message string) {
	macro.Rewrite("$_.message")
	return message
}

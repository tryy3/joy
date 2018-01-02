package dom

import "github.com/matthewmueller/joy/macro"

type DOMException struct {
}

func (*DOMException) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

func (*DOMException) Code() (code uint8) {
	macro.Rewrite("$_.code")
	return code
}

func (*DOMException) Message() (message string) {
	macro.Rewrite("$_.message")
	return message
}

func (*DOMException) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

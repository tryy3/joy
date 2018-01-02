package dom

import "github.com/matthewmueller/joy/macro"

type DOMError struct {
}

func (*DOMError) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

func (*DOMError) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

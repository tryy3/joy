package dom

import "github.com/matthewmueller/joy/macro"

type SVGNumber struct {
}

func (*SVGNumber) Value() (value float32) {
	macro.Rewrite("$_.value")
	return value
}

func (*SVGNumber) SetValue(value float32) {
	macro.Rewrite("$_.value = $1", value)
}

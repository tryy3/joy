package browser

import "github.com/matthewmueller/joy/macro"

type BarProp struct {
}

func (*BarProp) Visible() (visible bool) {
	macro.Rewrite("$_.visible")
	return visible
}

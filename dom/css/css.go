package css

import "github.com/matthewmueller/joy/macro"

type CSS struct {
}

func (*CSS) Supports(property string, value *string) (b bool) {
	macro.Rewrite("$_.supports($1, $2)", property, value)
	return b
}

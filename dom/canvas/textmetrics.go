package canvas

import "github.com/matthewmueller/joy/macro"

type TextMetrics struct {
}

func (*TextMetrics) Width() (width float32) {
	macro.Rewrite("$_.width")
	return width
}

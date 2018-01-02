package canvas

import "github.com/matthewmueller/joy/macro"

type CanvasGradient struct {
}

func (*CanvasGradient) AddColorStop(offset float32, color string) {
	macro.Rewrite("$_.addColorStop($1, $2)", offset, color)
}

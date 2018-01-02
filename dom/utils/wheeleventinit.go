package utils

import "github.com/matthewmueller/joy/dom/mouse"

type WheelEventInit struct {
	*mouse.MouseEventInit

	deltaMode	*uint
	deltaX		*float32
	deltaY		*float32
	deltaZ		*float32
}

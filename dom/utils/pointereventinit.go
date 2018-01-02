package utils

import "github.com/matthewmueller/joy/dom/mouse"

type PointerEventInit struct {
	*mouse.MouseEventInit

	height		*int
	isPrimary	*bool
	pointerId	*int
	pointerType	*string
	pressure	*float32
	tiltX		*int
	tiltY		*int
	width		*int
}

package webgl

import "github.com/matthewmueller/joy/dom/event"

type WebGLContextEventInit struct {
	*event.EventInit

	statusMessage *string
}

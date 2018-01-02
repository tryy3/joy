package device

import "github.com/matthewmueller/joy/dom/event"

type DeviceLightEventInit struct {
	*event.EventInit

	value	*float32
}

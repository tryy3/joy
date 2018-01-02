package device

import "github.com/matthewmueller/joy/dom/event"

type DeviceOrientationEventInit struct {
	*event.EventInit

	absolute	*bool
	alpha		*float32
	beta		*float32
	gamma		*float32
}

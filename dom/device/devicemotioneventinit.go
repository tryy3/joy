package device

import "github.com/matthewmueller/joy/dom/event"

type DeviceMotionEventInit struct {
	*event.EventInit

	acceleration                 *DeviceAccelerationDict
	accelerationIncludingGravity *DeviceAccelerationDict
	interval                     *float32
	rotationRate                 *DeviceRotationRateDict
}

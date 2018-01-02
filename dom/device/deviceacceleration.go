package device

import "github.com/matthewmueller/joy/macro"

type DeviceAcceleration struct {
}

func (*DeviceAcceleration) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*DeviceAcceleration) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*DeviceAcceleration) Z() (z float32) {
	macro.Rewrite("$_.z")
	return z
}

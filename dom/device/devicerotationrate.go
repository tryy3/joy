package device

import "github.com/matthewmueller/joy/macro"

type DeviceRotationRate struct {
}

func (*DeviceRotationRate) Alpha() (alpha float32) {
	macro.Rewrite("$_.alpha")
	return alpha
}

func (*DeviceRotationRate) Beta() (beta float32) {
	macro.Rewrite("$_.beta")
	return beta
}

func (*DeviceRotationRate) Gamma() (gamma float32) {
	macro.Rewrite("$_.gamma")
	return gamma
}

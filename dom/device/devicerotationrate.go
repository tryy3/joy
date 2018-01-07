package device

import "github.com/matthewmueller/joy/macro"

// DeviceRotationRate struct
// js:"DeviceRotationRate,omit"
type DeviceRotationRate struct {
}

// Alpha prop
// js:"alpha"
func (*DeviceRotationRate) Alpha() (alpha float32) {
	macro.Rewrite("$_.alpha")
	return alpha
}

// Beta prop
// js:"beta"
func (*DeviceRotationRate) Beta() (beta float32) {
	macro.Rewrite("$_.beta")
	return beta
}

// Gamma prop
// js:"gamma"
func (*DeviceRotationRate) Gamma() (gamma float32) {
	macro.Rewrite("$_.gamma")
	return gamma
}

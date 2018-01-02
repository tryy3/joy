package audio

import "github.com/matthewmueller/joy/macro"

type AudioListener struct {
}

func (*AudioListener) SetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) {
	macro.Rewrite("$_.setOrientation($1, $2, $3, $4, $5, $6)", x, y, z, xUp, yUp, zUp)
}

func (*AudioListener) SetPosition(x float32, y float32, z float32) {
	macro.Rewrite("$_.setPosition($1, $2, $3)", x, y, z)
}

func (*AudioListener) SetVelocity(x float32, y float32, z float32) {
	macro.Rewrite("$_.setVelocity($1, $2, $3)", x, y, z)
}

func (*AudioListener) DopplerFactor() (dopplerFactor float32) {
	macro.Rewrite("$_.dopplerFactor")
	return dopplerFactor
}

func (*AudioListener) SetDopplerFactor(dopplerFactor float32) {
	macro.Rewrite("$_.dopplerFactor = $1", dopplerFactor)
}

func (*AudioListener) SpeedOfSound() (speedOfSound float32) {
	macro.Rewrite("$_.speedOfSound")
	return speedOfSound
}

func (*AudioListener) SetSpeedOfSound(speedOfSound float32) {
	macro.Rewrite("$_.speedOfSound = $1", speedOfSound)
}

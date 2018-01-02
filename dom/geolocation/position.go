package geolocation

import (
	"github.com/matthewmueller/joy/macro"
)

type Position struct {
}

func (*Position) Coords() (coords *Coordinates) {
	macro.Rewrite("$_.coords")
	return coords
}

func (*Position) Timestamp() (timestamp int) {
	macro.Rewrite("$_.timestamp")
	return timestamp
}

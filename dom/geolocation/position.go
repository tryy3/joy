package geolocation

import "github.com/matthewmueller/joy/macro"

// Position struct
// js:"Position,omit"
type Position struct {
}

// Coords prop
// js:"coords"
func (*Position) Coords() (coords *Coordinates) {
	macro.Rewrite("$_.coords")
	return coords
}

// Timestamp prop
// js:"timestamp"
func (*Position) Timestamp() (timestamp int) {
	macro.Rewrite("$_.timestamp")
	return timestamp
}

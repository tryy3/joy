package geolocation

import "github.com/matthewmueller/joy/macro"

// Geolocation struct
// js:"Geolocation,omit"
type Geolocation struct {
}

// ClearWatch fn
// js:"clearWatch"
func (*Geolocation) ClearWatch(watchId int) {
	macro.Rewrite("$_.clearWatch($1)", watchId)
}

// GetCurrentPosition fn
// js:"getCurrentPosition"
func (*Geolocation) GetCurrentPosition(successCallback func(position *Position), errorCallback *func(err *PositionError), options *PositionOptions) {
	macro.Rewrite("$_.getCurrentPosition($1, $2, $3)", successCallback, errorCallback, options)
}

// WatchPosition fn
// js:"watchPosition"
func (*Geolocation) WatchPosition(successCallback func(position *Position), errorCallback *func(err *PositionError), options *PositionOptions) (i int) {
	macro.Rewrite("$_.watchPosition($1, $2, $3)", successCallback, errorCallback, options)
	return i
}

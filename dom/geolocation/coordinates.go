package geolocation

import "github.com/matthewmueller/joy/macro"

type Coordinates struct {
}

func (*Coordinates) Accuracy() (accuracy float32) {
	macro.Rewrite("$_.accuracy")
	return accuracy
}

func (*Coordinates) Altitude() (altitude float32) {
	macro.Rewrite("$_.altitude")
	return altitude
}

func (*Coordinates) AltitudeAccuracy() (altitudeAccuracy float32) {
	macro.Rewrite("$_.altitudeAccuracy")
	return altitudeAccuracy
}

func (*Coordinates) Heading() (heading float32) {
	macro.Rewrite("$_.heading")
	return heading
}

func (*Coordinates) Latitude() (latitude float32) {
	macro.Rewrite("$_.latitude")
	return latitude
}

func (*Coordinates) Longitude() (longitude float32) {
	macro.Rewrite("$_.longitude")
	return longitude
}

func (*Coordinates) Speed() (speed float32) {
	macro.Rewrite("$_.speed")
	return speed
}

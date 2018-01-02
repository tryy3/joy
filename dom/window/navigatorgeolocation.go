package window

import "github.com/matthewmueller/joy/dom/geolocation"

type NavigatorGeolocation interface {
	Geolocation() (geolocation *geolocation.Geolocation)
}

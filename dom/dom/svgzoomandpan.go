package dom

import "github.com/matthewmueller/joy/macro"

type SVGZoomAndPan struct {
}

func (*SVGZoomAndPan) ZoomAndPan() (zoomAndPan uint8) {
	macro.Rewrite("$_.zoomAndPan")
	return zoomAndPan
}

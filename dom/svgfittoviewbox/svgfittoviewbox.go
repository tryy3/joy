package svgfittoviewbox

import "github.com/matthewmueller/joy/dom/dom"

// SVGFitToViewBox interface
// js:"SVGFitToViewBox"
type SVGFitToViewBox interface {

	// PreserveAspectRatio prop
	// js:"preserveAspectRatio"
	// jsrewrite:"$_.preserveAspectRatio"
	PreserveAspectRatio() (preserveAspectRatio *dom.SVGAnimatedPreserveAspectRatio)

	// ViewBox prop
	// js:"viewBox"
	// jsrewrite:"$_.viewBox"
	ViewBox() (viewBox *dom.SVGAnimatedRect)
}

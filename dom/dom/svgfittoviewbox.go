package dom


import (
	"github.com/matthewmueller/joy/dom/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/joy/dom/svganimatedrect"
)

// SVGFitToViewBox interface
// js:"SVGFitToViewBox"
type SVGFitToViewBox interface {

	// PreserveAspectRatio prop
	// js:"preserveAspectRatio"
	// jsrewrite:"$_.preserveAspectRatio"
	PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio)

	// ViewBox prop
	// js:"viewBox"
	// jsrewrite:"$_.viewBox"
	ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect)
}

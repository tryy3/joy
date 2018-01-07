package svggradientelement

import (
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/svganimatedtransformlist"
	"github.com/matthewmueller/joy/dom/window"
)

// SVGGradientElement interface
// js:"SVGGradientElement"
type SVGGradientElement interface {
	window.SVGElement

	// Href prop
	// js:"href"
	// jsrewrite:"$_.href"
	Href() (href *dom.SVGAnimatedString)

	// GradientTransform prop
	// js:"gradientTransform"
	// jsrewrite:"$_.gradientTransform"
	GradientTransform() (gradientTransform *svganimatedtransformlist.SVGAnimatedTransformList)

	// GradientUnits prop
	// js:"gradientUnits"
	// jsrewrite:"$_.gradientUnits"
	GradientUnits() (gradientUnits *dom.SVGAnimatedEnumeration)

	// SpreadMethod prop
	// js:"spreadMethod"
	// jsrewrite:"$_.spreadMethod"
	SpreadMethod() (spreadMethod *dom.SVGAnimatedEnumeration)
}

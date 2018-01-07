package svgcomponenttransferfunctionelement

import (
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/svganimatednumberlist"
	"github.com/matthewmueller/joy/dom/window"
)

// SVGComponentTransferFunctionElement interface
// js:"SVGComponentTransferFunctionElement"
type SVGComponentTransferFunctionElement interface {
	window.SVGElement

	// Amplitude prop
	// js:"amplitude"
	// jsrewrite:"$_.amplitude"
	Amplitude() (amplitude *dom.SVGAnimatedNumber)

	// Exponent prop
	// js:"exponent"
	// jsrewrite:"$_.exponent"
	Exponent() (exponent *dom.SVGAnimatedNumber)

	// Intercept prop
	// js:"intercept"
	// jsrewrite:"$_.intercept"
	Intercept() (intercept *dom.SVGAnimatedNumber)

	// Offset prop
	// js:"offset"
	// jsrewrite:"$_.offset"
	Offset() (offset *dom.SVGAnimatedNumber)

	// Slope prop
	// js:"slope"
	// jsrewrite:"$_.slope"
	Slope() (slope *dom.SVGAnimatedNumber)

	// TableValues prop
	// js:"tableValues"
	// jsrewrite:"$_.tableValues"
	TableValues() (tableValues *svganimatednumberlist.SVGAnimatedNumberList)

	// Type prop
	// js:"type"
	// jsrewrite:"$_.type"
	Type() (kind *dom.SVGAnimatedEnumeration)
}

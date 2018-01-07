package svgfilterprimitivestandardattributes

import "github.com/matthewmueller/joy/dom/dom"

// SVGFilterPrimitiveStandardAttributes interface
// js:"SVGFilterPrimitiveStandardAttributes"
type SVGFilterPrimitiveStandardAttributes interface {

	// Height prop
	// js:"height"
	// jsrewrite:"$_.height"
	Height() (height *dom.SVGAnimatedLength)

	// Result prop
	// js:"result"
	// jsrewrite:"$_.result"
	Result() (result *dom.SVGAnimatedString)

	// Width prop
	// js:"width"
	// jsrewrite:"$_.width"
	Width() (width *dom.SVGAnimatedLength)

	// X prop
	// js:"x"
	// jsrewrite:"$_.x"
	X() (x *dom.SVGAnimatedLength)

	// Y prop
	// js:"y"
	// jsrewrite:"$_.y"
	Y() (y *dom.SVGAnimatedLength)
}

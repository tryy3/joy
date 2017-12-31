package dom


import "github.com/matthewmueller/joy/dom/svganimatedstring"

// SVGURIReference interface
// js:"SVGURIReference"
type SVGURIReference interface {

	// Href prop
	// js:"href"
	// jsrewrite:"$_.href"
	Href() (href *svganimatedstring.SVGAnimatedString)
}

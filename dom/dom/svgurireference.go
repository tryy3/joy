package dom

// SVGURIReference interface
// js:"SVGURIReference"
type SVGURIReference interface {

	// Href prop
	// js:"href"
	// jsrewrite:"$_.href"
	Href() (href *SVGAnimatedString)
}

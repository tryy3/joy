package dom

import "github.com/matthewmueller/joy/macro"

// SVGAnimatedLength struct
// js:"SVGAnimatedLength,omit"
type SVGAnimatedLength struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedLength) AnimVal() (animVal *SVGLength) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedLength) BaseVal() (baseVal *SVGLength) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

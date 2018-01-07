package dom

import "github.com/matthewmueller/joy/macro"

// SVGAnimatedPreserveAspectRatio struct
// js:"SVGAnimatedPreserveAspectRatio,omit"
type SVGAnimatedPreserveAspectRatio struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedPreserveAspectRatio) AnimVal() (animVal *SVGPreserveAspectRatio) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedPreserveAspectRatio) BaseVal() (baseVal *SVGPreserveAspectRatio) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

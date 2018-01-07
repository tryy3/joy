package dom

import "github.com/matthewmueller/joy/macro"

// SVGAnimatedRect struct
// js:"SVGAnimatedRect,omit"
type SVGAnimatedRect struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedRect) AnimVal() (animVal *SVGRect) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedRect) BaseVal() (baseVal *SVGRect) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

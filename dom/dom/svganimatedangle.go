package dom

import "github.com/matthewmueller/joy/macro"

// SVGAnimatedAngle struct
// js:"SVGAnimatedAngle,omit"
type SVGAnimatedAngle struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedAngle) AnimVal() (animVal *SVGAngle) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedAngle) BaseVal() (baseVal *SVGAngle) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

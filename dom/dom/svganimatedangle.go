package dom


import (
	"github.com/matthewmueller/joy/dom/svgangle"
	"github.com/matthewmueller/joy/macro"
)

// SVGAnimatedAngle struct
// js:"SVGAnimatedAngle,omit"
type SVGAnimatedAngle struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedAngle) AnimVal() (animVal *svgangle.SVGAngle) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedAngle) BaseVal() (baseVal *svgangle.SVGAngle) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

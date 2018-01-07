package svganimatedlengthlist

import (
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/macro"
)

// SVGAnimatedLengthList struct
// js:"SVGAnimatedLengthList,omit"
type SVGAnimatedLengthList struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedLengthList) AnimVal() (animVal *dom.SVGLengthList) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedLengthList) BaseVal() (baseVal *dom.SVGLengthList) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

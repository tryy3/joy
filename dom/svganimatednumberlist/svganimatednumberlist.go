package svganimatednumberlist

import (
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/macro"
)

// SVGAnimatedNumberList struct
// js:"SVGAnimatedNumberList,omit"
type SVGAnimatedNumberList struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedNumberList) AnimVal() (animVal *dom.SVGNumberList) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedNumberList) BaseVal() (baseVal *dom.SVGNumberList) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

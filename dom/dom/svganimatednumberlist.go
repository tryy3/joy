package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGAnimatedNumberList struct {
}

func (*SVGAnimatedNumberList) AnimVal() (animVal *SVGNumberList) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedNumberList) BaseVal() (baseVal *SVGNumberList) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGAnimatedLengthList struct {
}

func (*SVGAnimatedLengthList) AnimVal() (animVal *SVGLengthList) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedLengthList) BaseVal() (baseVal *SVGLengthList) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

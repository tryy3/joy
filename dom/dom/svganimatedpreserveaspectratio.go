package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGAnimatedPreserveAspectRatio struct {
}

func (*SVGAnimatedPreserveAspectRatio) AnimVal() (animVal *SVGPreserveAspectRatio) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedPreserveAspectRatio) BaseVal() (baseVal *SVGPreserveAspectRatio) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

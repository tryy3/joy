package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGAnimatedTransformList struct {
}

func (*SVGAnimatedTransformList) AnimVal() (animVal *SVGTransformList) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedTransformList) BaseVal() (baseVal *SVGTransformList) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

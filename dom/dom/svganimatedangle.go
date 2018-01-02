package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGAnimatedAngle struct {
}

func (*SVGAnimatedAngle) AnimVal() (animVal *SVGAngle) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedAngle) BaseVal() (baseVal *SVGAngle) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

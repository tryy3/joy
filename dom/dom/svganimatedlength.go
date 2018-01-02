package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGAnimatedLength struct {
}

func (*SVGAnimatedLength) AnimVal() (animVal *SVGLength) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedLength) BaseVal() (baseVal *SVGLength) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

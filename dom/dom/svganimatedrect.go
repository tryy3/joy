package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGAnimatedRect struct {
}

func (*SVGAnimatedRect) AnimVal() (animVal *SVGRect) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedRect) BaseVal() (baseVal *SVGRect) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

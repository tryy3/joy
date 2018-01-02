package dom

import "github.com/matthewmueller/joy/macro"

type SVGAnimatedBoolean struct {
}

func (*SVGAnimatedBoolean) AnimVal() (animVal bool) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedBoolean) BaseVal() (baseVal bool) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

func (*SVGAnimatedBoolean) SetBaseVal(baseVal bool) {
	macro.Rewrite("$_.baseVal = $1", baseVal)
}

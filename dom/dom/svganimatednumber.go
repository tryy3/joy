package dom

import "github.com/matthewmueller/joy/macro"

type SVGAnimatedNumber struct {
}

func (*SVGAnimatedNumber) AnimVal() (animVal float32) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedNumber) BaseVal() (baseVal float32) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

func (*SVGAnimatedNumber) SetBaseVal(baseVal float32) {
	macro.Rewrite("$_.baseVal = $1", baseVal)
}

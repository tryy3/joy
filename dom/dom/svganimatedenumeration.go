package dom

import "github.com/matthewmueller/joy/macro"

type SVGAnimatedEnumeration struct {
}

func (*SVGAnimatedEnumeration) AnimVal() (animVal uint8) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedEnumeration) BaseVal() (baseVal uint8) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

func (*SVGAnimatedEnumeration) SetBaseVal(baseVal uint8) {
	macro.Rewrite("$_.baseVal = $1", baseVal)
}

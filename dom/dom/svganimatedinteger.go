package dom

import "github.com/matthewmueller/joy/macro"

type SVGAnimatedInteger struct {
}

func (*SVGAnimatedInteger) AnimVal() (animVal int) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedInteger) BaseVal() (baseVal int) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

func (*SVGAnimatedInteger) SetBaseVal(baseVal int) {
	macro.Rewrite("$_.baseVal = $1", baseVal)
}

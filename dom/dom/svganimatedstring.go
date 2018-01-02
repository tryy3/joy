package dom

import "github.com/matthewmueller/joy/macro"

type SVGAnimatedString struct {
}

func (*SVGAnimatedString) AnimVal() (animVal string) {
	macro.Rewrite("$_.animVal")
	return animVal
}

func (*SVGAnimatedString) BaseVal() (baseVal string) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

func (*SVGAnimatedString) SetBaseVal(baseVal string) {
	macro.Rewrite("$_.baseVal = $1", baseVal)
}

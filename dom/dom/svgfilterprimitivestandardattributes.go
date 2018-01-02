package dom

type SVGFilterPrimitiveStandardAttributes interface {
	Height() (height *SVGAnimatedLength)

	Result() (result *SVGAnimatedString)

	Width() (width *SVGAnimatedLength)

	X() (x *SVGAnimatedLength)

	Y() (y *SVGAnimatedLength)
}

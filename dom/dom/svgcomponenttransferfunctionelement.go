package dom

type SVGComponentTransferFunctionElement interface {
	SVGElement

	Amplitude() (amplitude *SVGAnimatedNumber)

	Exponent() (exponent *SVGAnimatedNumber)

	Intercept() (intercept *SVGAnimatedNumber)

	Offset() (offset *SVGAnimatedNumber)

	Slope() (slope *SVGAnimatedNumber)

	TableValues() (tableValues *SVGAnimatedNumberList)

	Type() (kind *SVGAnimatedEnumeration)
}

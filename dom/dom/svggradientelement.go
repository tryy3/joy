package dom

type SVGGradientElement interface {
	SVGElement

	Href() (href *SVGAnimatedString)

	GradientTransform() (gradientTransform *SVGAnimatedTransformList)

	GradientUnits() (gradientUnits *SVGAnimatedEnumeration)

	SpreadMethod() (spreadMethod *SVGAnimatedEnumeration)
}

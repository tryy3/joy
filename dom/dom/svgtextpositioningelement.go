package dom

type SVGTextPositioningElement interface {
	SVGTextContentElement

	Dx() (dx *SVGAnimatedLengthList)

	Dy() (dy *SVGAnimatedLengthList)

	Rotate() (rotate *SVGAnimatedNumberList)

	X() (x *SVGAnimatedLengthList)

	Y() (y *SVGAnimatedLengthList)
}

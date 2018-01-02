package dom

type SVGFitToViewBox interface {
	PreserveAspectRatio() (preserveAspectRatio *SVGAnimatedPreserveAspectRatio)

	ViewBox() (viewBox *SVGAnimatedRect)
}

package dom

type SVGAnimatedPoints interface {
	AnimatedPoints() (animatedPoints *SVGPointList)

	Points() (points *SVGPointList)
}

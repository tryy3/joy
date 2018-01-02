package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegCurvetoQuadraticSmoothAbs)(nil)

type SVGPathSegCurvetoQuadraticSmoothAbs struct {
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

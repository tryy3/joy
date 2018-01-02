package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegCurvetoQuadraticAbs)(nil)

type SVGPathSegCurvetoQuadraticAbs struct {
}

func (*SVGPathSegCurvetoQuadraticAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticAbs) X1() (x1 float32) {
	macro.Rewrite("$_.x1")
	return x1
}

func (*SVGPathSegCurvetoQuadraticAbs) SetX1(x1 float32) {
	macro.Rewrite("$_.x1 = $1", x1)
}

func (*SVGPathSegCurvetoQuadraticAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegCurvetoQuadraticAbs) Y1() (y1 float32) {
	macro.Rewrite("$_.y1")
	return y1
}

func (*SVGPathSegCurvetoQuadraticAbs) SetY1(y1 float32) {
	macro.Rewrite("$_.y1 = $1", y1)
}

func (*SVGPathSegCurvetoQuadraticAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegCurvetoQuadraticAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegCurvetoQuadraticRel)(nil)

type SVGPathSegCurvetoQuadraticRel struct {
}

func (*SVGPathSegCurvetoQuadraticRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticRel) X1() (x1 float32) {
	macro.Rewrite("$_.x1")
	return x1
}

func (*SVGPathSegCurvetoQuadraticRel) SetX1(x1 float32) {
	macro.Rewrite("$_.x1 = $1", x1)
}

func (*SVGPathSegCurvetoQuadraticRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegCurvetoQuadraticRel) Y1() (y1 float32) {
	macro.Rewrite("$_.y1")
	return y1
}

func (*SVGPathSegCurvetoQuadraticRel) SetY1(y1 float32) {
	macro.Rewrite("$_.y1 = $1", y1)
}

func (*SVGPathSegCurvetoQuadraticRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegCurvetoQuadraticRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

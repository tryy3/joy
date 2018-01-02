package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegCurvetoCubicSmoothRel)(nil)

type SVGPathSegCurvetoCubicSmoothRel struct {
}

func (*SVGPathSegCurvetoCubicSmoothRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegCurvetoCubicSmoothRel) X2() (x2 float32) {
	macro.Rewrite("$_.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetX2(x2 float32) {
	macro.Rewrite("$_.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicSmoothRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegCurvetoCubicSmoothRel) Y2() (y2 float32) {
	macro.Rewrite("$_.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetY2(y2 float32) {
	macro.Rewrite("$_.y2 = $1", y2)
}

func (*SVGPathSegCurvetoCubicSmoothRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegCurvetoCubicSmoothRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

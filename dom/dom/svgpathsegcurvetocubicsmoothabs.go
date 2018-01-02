package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegCurvetoCubicSmoothAbs)(nil)

type SVGPathSegCurvetoCubicSmoothAbs struct {
}

func (*SVGPathSegCurvetoCubicSmoothAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegCurvetoCubicSmoothAbs) X2() (x2 float32) {
	macro.Rewrite("$_.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetX2(x2 float32) {
	macro.Rewrite("$_.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicSmoothAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegCurvetoCubicSmoothAbs) Y2() (y2 float32) {
	macro.Rewrite("$_.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetY2(y2 float32) {
	macro.Rewrite("$_.y2 = $1", y2)
}

func (*SVGPathSegCurvetoCubicSmoothAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegCurvetoCubicSmoothAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

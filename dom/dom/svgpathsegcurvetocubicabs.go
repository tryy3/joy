package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegCurvetoCubicAbs)(nil)

type SVGPathSegCurvetoCubicAbs struct {
}

func (*SVGPathSegCurvetoCubicAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegCurvetoCubicAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegCurvetoCubicAbs) X1() (x1 float32) {
	macro.Rewrite("$_.x1")
	return x1
}

func (*SVGPathSegCurvetoCubicAbs) SetX1(x1 float32) {
	macro.Rewrite("$_.x1 = $1", x1)
}

func (*SVGPathSegCurvetoCubicAbs) X2() (x2 float32) {
	macro.Rewrite("$_.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicAbs) SetX2(x2 float32) {
	macro.Rewrite("$_.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegCurvetoCubicAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegCurvetoCubicAbs) Y1() (y1 float32) {
	macro.Rewrite("$_.y1")
	return y1
}

func (*SVGPathSegCurvetoCubicAbs) SetY1(y1 float32) {
	macro.Rewrite("$_.y1 = $1", y1)
}

func (*SVGPathSegCurvetoCubicAbs) Y2() (y2 float32) {
	macro.Rewrite("$_.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicAbs) SetY2(y2 float32) {
	macro.Rewrite("$_.y2 = $1", y2)
}

func (*SVGPathSegCurvetoCubicAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegCurvetoCubicAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

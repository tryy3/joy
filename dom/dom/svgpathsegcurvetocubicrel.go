package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegCurvetoCubicRel)(nil)

type SVGPathSegCurvetoCubicRel struct {
}

func (*SVGPathSegCurvetoCubicRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegCurvetoCubicRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegCurvetoCubicRel) X1() (x1 float32) {
	macro.Rewrite("$_.x1")
	return x1
}

func (*SVGPathSegCurvetoCubicRel) SetX1(x1 float32) {
	macro.Rewrite("$_.x1 = $1", x1)
}

func (*SVGPathSegCurvetoCubicRel) X2() (x2 float32) {
	macro.Rewrite("$_.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicRel) SetX2(x2 float32) {
	macro.Rewrite("$_.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegCurvetoCubicRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegCurvetoCubicRel) Y1() (y1 float32) {
	macro.Rewrite("$_.y1")
	return y1
}

func (*SVGPathSegCurvetoCubicRel) SetY1(y1 float32) {
	macro.Rewrite("$_.y1 = $1", y1)
}

func (*SVGPathSegCurvetoCubicRel) Y2() (y2 float32) {
	macro.Rewrite("$_.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicRel) SetY2(y2 float32) {
	macro.Rewrite("$_.y2 = $1", y2)
}

func (*SVGPathSegCurvetoCubicRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegCurvetoCubicRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegMovetoRel)(nil)

type SVGPathSegMovetoRel struct {
}

func (*SVGPathSegMovetoRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegMovetoRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegMovetoRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegMovetoRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegMovetoRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegMovetoRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

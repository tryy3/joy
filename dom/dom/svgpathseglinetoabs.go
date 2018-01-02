package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegLinetoAbs)(nil)

type SVGPathSegLinetoAbs struct {
}

func (*SVGPathSegLinetoAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegLinetoAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegLinetoAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegLinetoAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegLinetoAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegLinetoAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

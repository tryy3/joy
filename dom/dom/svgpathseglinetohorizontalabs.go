package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegLinetoHorizontalAbs)(nil)

type SVGPathSegLinetoHorizontalAbs struct {
}

func (*SVGPathSegLinetoHorizontalAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegLinetoHorizontalAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegLinetoHorizontalAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegLinetoHorizontalAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

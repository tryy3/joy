package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegLinetoVerticalAbs)(nil)

type SVGPathSegLinetoVerticalAbs struct {
}

func (*SVGPathSegLinetoVerticalAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegLinetoVerticalAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegLinetoVerticalAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegLinetoVerticalAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

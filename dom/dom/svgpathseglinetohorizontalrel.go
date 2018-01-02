package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegLinetoHorizontalRel)(nil)

type SVGPathSegLinetoHorizontalRel struct {
}

func (*SVGPathSegLinetoHorizontalRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegLinetoHorizontalRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegLinetoHorizontalRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegLinetoHorizontalRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

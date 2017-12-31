package dom


import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegCurvetoQuadraticAbs)(nil)

// SVGPathSegCurvetoQuadraticAbs struct
// js:"SVGPathSegCurvetoQuadraticAbs,omit"
type SVGPathSegCurvetoQuadraticAbs struct {
}

// X prop
// js:"x"
func (*SVGPathSegCurvetoQuadraticAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegCurvetoQuadraticAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// X1 prop
// js:"x1"
func (*SVGPathSegCurvetoQuadraticAbs) X1() (x1 float32) {
	macro.Rewrite("$_.x1")
	return x1
}

// SetX1 prop
// js:"x1"
func (*SVGPathSegCurvetoQuadraticAbs) SetX1(x1 float32) {
	macro.Rewrite("$_.x1 = $1", x1)
}

// Y prop
// js:"y"
func (*SVGPathSegCurvetoQuadraticAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegCurvetoQuadraticAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// Y1 prop
// js:"y1"
func (*SVGPathSegCurvetoQuadraticAbs) Y1() (y1 float32) {
	macro.Rewrite("$_.y1")
	return y1
}

// SetY1 prop
// js:"y1"
func (*SVGPathSegCurvetoQuadraticAbs) SetY1(y1 float32) {
	macro.Rewrite("$_.y1 = $1", y1)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegCurvetoQuadraticAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegCurvetoQuadraticAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

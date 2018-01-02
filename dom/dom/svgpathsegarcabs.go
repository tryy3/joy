package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegArcAbs)(nil)

type SVGPathSegArcAbs struct {
}

func (*SVGPathSegArcAbs) Angle() (angle float32) {
	macro.Rewrite("$_.angle")
	return angle
}

func (*SVGPathSegArcAbs) SetAngle(angle float32) {
	macro.Rewrite("$_.angle = $1", angle)
}

func (*SVGPathSegArcAbs) LargeArcFlag() (largeArcFlag bool) {
	macro.Rewrite("$_.largeArcFlag")
	return largeArcFlag
}

func (*SVGPathSegArcAbs) SetLargeArcFlag(largeArcFlag bool) {
	macro.Rewrite("$_.largeArcFlag = $1", largeArcFlag)
}

func (*SVGPathSegArcAbs) R1() (r1 float32) {
	macro.Rewrite("$_.r1")
	return r1
}

func (*SVGPathSegArcAbs) SetR1(r1 float32) {
	macro.Rewrite("$_.r1 = $1", r1)
}

func (*SVGPathSegArcAbs) R2() (r2 float32) {
	macro.Rewrite("$_.r2")
	return r2
}

func (*SVGPathSegArcAbs) SetR2(r2 float32) {
	macro.Rewrite("$_.r2 = $1", r2)
}

func (*SVGPathSegArcAbs) SweepFlag() (sweepFlag bool) {
	macro.Rewrite("$_.sweepFlag")
	return sweepFlag
}

func (*SVGPathSegArcAbs) SetSweepFlag(sweepFlag bool) {
	macro.Rewrite("$_.sweepFlag = $1", sweepFlag)
}

func (*SVGPathSegArcAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegArcAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegArcAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegArcAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegArcAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegArcAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

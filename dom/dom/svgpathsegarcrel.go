package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ SVGPathSeg = (*SVGPathSegArcRel)(nil)

type SVGPathSegArcRel struct {
}

func (*SVGPathSegArcRel) Angle() (angle float32) {
	macro.Rewrite("$_.angle")
	return angle
}

func (*SVGPathSegArcRel) SetAngle(angle float32) {
	macro.Rewrite("$_.angle = $1", angle)
}

func (*SVGPathSegArcRel) LargeArcFlag() (largeArcFlag bool) {
	macro.Rewrite("$_.largeArcFlag")
	return largeArcFlag
}

func (*SVGPathSegArcRel) SetLargeArcFlag(largeArcFlag bool) {
	macro.Rewrite("$_.largeArcFlag = $1", largeArcFlag)
}

func (*SVGPathSegArcRel) R1() (r1 float32) {
	macro.Rewrite("$_.r1")
	return r1
}

func (*SVGPathSegArcRel) SetR1(r1 float32) {
	macro.Rewrite("$_.r1 = $1", r1)
}

func (*SVGPathSegArcRel) R2() (r2 float32) {
	macro.Rewrite("$_.r2")
	return r2
}

func (*SVGPathSegArcRel) SetR2(r2 float32) {
	macro.Rewrite("$_.r2 = $1", r2)
}

func (*SVGPathSegArcRel) SweepFlag() (sweepFlag bool) {
	macro.Rewrite("$_.sweepFlag")
	return sweepFlag
}

func (*SVGPathSegArcRel) SetSweepFlag(sweepFlag bool) {
	macro.Rewrite("$_.sweepFlag = $1", sweepFlag)
}

func (*SVGPathSegArcRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPathSegArcRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPathSegArcRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPathSegArcRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

func (*SVGPathSegArcRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

func (*SVGPathSegArcRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

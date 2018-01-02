package canvas

import "github.com/matthewmueller/joy/macro"

func New(path *Path2d) *Path2d {
	macro.Rewrite("new Path2D($1)", path)
	return &Path2d{}
}

type Path2d struct {
}

func (*Path2d) Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	macro.Rewrite("$_.arc($1, $2, $3, $4, $5, $6)", x, y, radius, startAngle, endAngle, anticlockwise)
}

func (*Path2d) ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32) {
	macro.Rewrite("$_.arcTo($1, $2, $3, $4, $5)", x1, y1, x2, y2, radius)
}

func (*Path2d) BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32) {
	macro.Rewrite("$_.bezierCurveTo($1, $2, $3, $4, $5, $6)", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (*Path2d) ClosePath() {
	macro.Rewrite("$_.closePath()")
}

func (*Path2d) Ellipse(x float32, y float32, radiusX float32, radiusY float32, rotation float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	macro.Rewrite("$_.ellipse($1, $2, $3, $4, $5, $6, $7, $8)", x, y, radiusX, radiusY, rotation, startAngle, endAngle, anticlockwise)
}

func (*Path2d) LineTo(x float32, y float32) {
	macro.Rewrite("$_.lineTo($1, $2)", x, y)
}

func (*Path2d) MoveTo(x float32, y float32) {
	macro.Rewrite("$_.moveTo($1, $2)", x, y)
}

func (*Path2d) QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32) {
	macro.Rewrite("$_.quadraticCurveTo($1, $2, $3, $4)", cpx, cpy, x, y)
}

func (*Path2d) Rect(x float32, y float32, w float32, h float32) {
	macro.Rewrite("$_.rect($1, $2, $3, $4)", x, y, w, h)
}

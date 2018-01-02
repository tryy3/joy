package canvas

type CanvasPathMethods interface {
	Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise *bool)

	ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32)

	BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32)

	ClosePath()

	Ellipse(x float32, y float32, radiusX float32, radiusY float32, rotation float32, startAngle float32, endAngle float32, anticlockwise *bool)

	LineTo(x float32, y float32)

	MoveTo(x float32, y float32)

	QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32)

	Rect(x float32, y float32, w float32, h float32)
}

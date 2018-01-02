package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGPoint struct {
}

func (*SVGPoint) MatrixTransform(matrix *SVGMatrix) (s *SVGPoint) {
	macro.Rewrite("$_.matrixTransform($1)", matrix)
	return s
}

func (*SVGPoint) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGPoint) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGPoint) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGPoint) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

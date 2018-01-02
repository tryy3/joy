package dom

import (
	"github.com/matthewmueller/joy/macro"
)

type SVGTransform struct {
}

func (*SVGTransform) SetMatrix(matrix *SVGMatrix) {
	macro.Rewrite("$_.setMatrix($1)", matrix)
}

func (*SVGTransform) SetRotate(angle float32, cx float32, cy float32) {
	macro.Rewrite("$_.setRotate($1, $2, $3)", angle, cx, cy)
}

func (*SVGTransform) SetScale(sx float32, sy float32) {
	macro.Rewrite("$_.setScale($1, $2)", sx, sy)
}

func (*SVGTransform) SetSkewX(angle float32) {
	macro.Rewrite("$_.setSkewX($1)", angle)
}

func (*SVGTransform) SetSkewY(angle float32) {
	macro.Rewrite("$_.setSkewY($1)", angle)
}

func (*SVGTransform) SetTranslate(tx float32, ty float32) {
	macro.Rewrite("$_.setTranslate($1, $2)", tx, ty)
}

func (*SVGTransform) Angle() (angle float32) {
	macro.Rewrite("$_.angle")
	return angle
}

func (*SVGTransform) Matrix() (matrix *SVGMatrix) {
	macro.Rewrite("$_.matrix")
	return matrix
}

func (*SVGTransform) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}

package dom

import "github.com/matthewmueller/joy/macro"

type SVGMatrix struct {
}

func (*SVGMatrix) FlipX() (s *SVGMatrix) {
	macro.Rewrite("$_.flipX()")
	return s
}

func (*SVGMatrix) FlipY() (s *SVGMatrix) {
	macro.Rewrite("$_.flipY()")
	return s
}

func (*SVGMatrix) Inverse() (s *SVGMatrix) {
	macro.Rewrite("$_.inverse()")
	return s
}

func (*SVGMatrix) Multiply(secondMatrix *SVGMatrix) (s *SVGMatrix) {
	macro.Rewrite("$_.multiply($1)", secondMatrix)
	return s
}

func (*SVGMatrix) Rotate(angle float32) (s *SVGMatrix) {
	macro.Rewrite("$_.rotate($1)", angle)
	return s
}

func (*SVGMatrix) RotateFromVector(x float32, y float32) (s *SVGMatrix) {
	macro.Rewrite("$_.rotateFromVector($1, $2)", x, y)
	return s
}

func (*SVGMatrix) Scale(scaleFactor float32) (s *SVGMatrix) {
	macro.Rewrite("$_.scale($1)", scaleFactor)
	return s
}

func (*SVGMatrix) ScaleNonUniform(scaleFactorX float32, scaleFactorY float32) (s *SVGMatrix) {
	macro.Rewrite("$_.scaleNonUniform($1, $2)", scaleFactorX, scaleFactorY)
	return s
}

func (*SVGMatrix) SkewX(angle float32) (s *SVGMatrix) {
	macro.Rewrite("$_.skewX($1)", angle)
	return s
}

func (*SVGMatrix) SkewY(angle float32) (s *SVGMatrix) {
	macro.Rewrite("$_.skewY($1)", angle)
	return s
}

func (*SVGMatrix) Translate(x float32, y float32) (s *SVGMatrix) {
	macro.Rewrite("$_.translate($1, $2)", x, y)
	return s
}

func (*SVGMatrix) A() (a float32) {
	macro.Rewrite("$_.a")
	return a
}

func (*SVGMatrix) SetA(a float32) {
	macro.Rewrite("$_.a = $1", a)
}

func (*SVGMatrix) B() (b float32) {
	macro.Rewrite("$_.b")
	return b
}

func (*SVGMatrix) SetB(b float32) {
	macro.Rewrite("$_.b = $1", b)
}

func (*SVGMatrix) C() (c float32) {
	macro.Rewrite("$_.c")
	return c
}

func (*SVGMatrix) SetC(c float32) {
	macro.Rewrite("$_.c = $1", c)
}

func (*SVGMatrix) D() (d float32) {
	macro.Rewrite("$_.d")
	return d
}

func (*SVGMatrix) SetD(d float32) {
	macro.Rewrite("$_.d = $1", d)
}

func (*SVGMatrix) E() (e float32) {
	macro.Rewrite("$_.e")
	return e
}

func (*SVGMatrix) SetE(e float32) {
	macro.Rewrite("$_.e = $1", e)
}

func (*SVGMatrix) F() (f float32) {
	macro.Rewrite("$_.f")
	return f
}

func (*SVGMatrix) SetF(f float32) {
	macro.Rewrite("$_.f = $1", f)
}

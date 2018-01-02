package webgl

import "github.com/matthewmueller/joy/macro"

type WebGLShaderPrecisionFormat struct {
}

func (*WebGLShaderPrecisionFormat) Precision() (precision int) {
	macro.Rewrite("$_.precision")
	return precision
}

func (*WebGLShaderPrecisionFormat) RangeMax() (rangeMax int) {
	macro.Rewrite("$_.rangeMax")
	return rangeMax
}

func (*WebGLShaderPrecisionFormat) RangeMin() (rangeMin int) {
	macro.Rewrite("$_.rangeMin")
	return rangeMin
}

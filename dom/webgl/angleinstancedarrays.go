package webgl

import "github.com/matthewmueller/joy/macro"

type ANGLEInstancedArrays struct {
}

func (*ANGLEInstancedArrays) DrawArraysInstancedANGLE(mode uint, first int, count int, primcount int) {
	macro.Rewrite("$_.drawArraysInstancedANGLE($1, $2, $3, $4)", mode, first, count, primcount)
}

func (*ANGLEInstancedArrays) DrawElementsInstancedANGLE(mode uint, count int, kind uint, offset int64, primcount int) {
	macro.Rewrite("$_.drawElementsInstancedANGLE($1, $2, $3, $4, $5)", mode, count, kind, offset, primcount)
}

func (*ANGLEInstancedArrays) VertexAttribDivisorANGLE(index uint, divisor uint) {
	macro.Rewrite("$_.vertexAttribDivisorANGLE($1, $2)", index, divisor)
}

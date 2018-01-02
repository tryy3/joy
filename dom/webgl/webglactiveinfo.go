package webgl

import "github.com/matthewmueller/joy/macro"

type WebGLActiveInfo struct {
}

func (*WebGLActiveInfo) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*WebGLActiveInfo) Size() (size int) {
	macro.Rewrite("$_.size")
	return size
}

func (*WebGLActiveInfo) Type() (kind uint) {
	macro.Rewrite("$_.type")
	return kind
}

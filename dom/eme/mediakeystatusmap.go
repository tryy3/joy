package eme

import (
	"github.com/matthewmueller/joy/macro"
)

type MediaKeyStatusMap struct {
}

func (*MediaKeyStatusMap) ForEach(callback func(keyId []byte, status *MediaKeyStatus)) {
	macro.Rewrite("$_.forEach($1)", callback)
}

func (*MediaKeyStatusMap) Get(keyId []byte) (m *MediaKeyStatus) {
	macro.Rewrite("$_.get($1)", keyId)
	return m
}

func (*MediaKeyStatusMap) Has(keyId []byte) (b bool) {
	macro.Rewrite("$_.has($1)", keyId)
	return b
}

func (*MediaKeyStatusMap) Size() (size uint) {
	macro.Rewrite("$_.size")
	return size
}

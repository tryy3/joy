package crypto

import (
	"github.com/matthewmueller/joy/macro"
)

type RandomSource interface {
	GetRandomValues(array []byte) (b []byte)
}

type KeyAlgorithm struct {
	name *string
}

type Crypto struct {
}

func (*Crypto) GetRandomValues(array []byte) (b []byte) {
	macro.Rewrite("$_.getRandomValues($1)", array)
	return b
}

func (*Crypto) Subtle() (subtle *SubtleCrypto) {
	macro.Rewrite("$_.subtle")
	return subtle
}

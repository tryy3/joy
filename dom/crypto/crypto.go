package crypto


import (
	"github.com/matthewmueller/joy/macro"
)

// RandomSource interface
// js:"RandomSource"
type RandomSource interface {

	// GetRandomValues
	// js:"getRandomValues"
	// jsrewrite:"$_.getRandomValues($1)"
	GetRandomValues(array []byte) (b []byte)
}

type KeyAlgorithm struct {
	name *string
}

// Crypto struct
// js:"Crypto,omit"
type Crypto struct {
}

// GetRandomValues fn
// js:"getRandomValues"
func (*Crypto) GetRandomValues(array []byte) (b []byte) {
	macro.Rewrite("$_.getRandomValues($1)", array)
	return b
}

// Subtle prop
// js:"subtle"
func (*Crypto) Subtle() (subtle *SubtleCrypto) {
	macro.Rewrite("$_.subtle")
	return subtle
}

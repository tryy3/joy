package crypto

import (
	"github.com/matthewmueller/joy/macro"
)

type CryptoKey struct {
}

func (*CryptoKey) Algorithm() (algorithm *KeyAlgorithm) {
	macro.Rewrite("$_.algorithm")
	return algorithm
}

func (*CryptoKey) Extractable() (extractable bool) {
	macro.Rewrite("$_.extractable")
	return extractable
}

func (*CryptoKey) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

func (*CryptoKey) Usages() (usages []string) {
	macro.Rewrite("$_.usages")
	return usages
}

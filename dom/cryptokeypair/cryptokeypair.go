package cryptokeypair

import (
	"github.com/matthewmueller/joy/dom/crypto"
	"github.com/matthewmueller/joy/macro"
)

// CryptoKeyPair struct
// js:"CryptoKeyPair,omit"
type CryptoKeyPair struct {
}

// PrivateKey prop
// js:"privateKey"
func (*CryptoKeyPair) PrivateKey() (privateKey *crypto.CryptoKey) {
	macro.Rewrite("$_.privateKey")
	return privateKey
}

// SetPrivateKey prop
// js:"privateKey"
func (*CryptoKeyPair) SetPrivateKey(privateKey *crypto.CryptoKey) {
	macro.Rewrite("$_.privateKey = $1", privateKey)
}

// PublicKey prop
// js:"publicKey"
func (*CryptoKeyPair) PublicKey() (publicKey *crypto.CryptoKey) {
	macro.Rewrite("$_.publicKey")
	return publicKey
}

// SetPublicKey prop
// js:"publicKey"
func (*CryptoKeyPair) SetPublicKey(publicKey *crypto.CryptoKey) {
	macro.Rewrite("$_.publicKey = $1", publicKey)
}

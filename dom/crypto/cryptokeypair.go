package crypto

import (
	"github.com/matthewmueller/joy/macro"
)

type CryptoKeyPair struct {
}

func (*CryptoKeyPair) PrivateKey() (privateKey *CryptoKey) {
	macro.Rewrite("$_.privateKey")
	return privateKey
}

func (*CryptoKeyPair) SetPrivateKey(privateKey *CryptoKey) {
	macro.Rewrite("$_.privateKey = $1", privateKey)
}

func (*CryptoKeyPair) PublicKey() (publicKey *CryptoKey) {
	macro.Rewrite("$_.publicKey")
	return publicKey
}

func (*CryptoKeyPair) SetPublicKey(publicKey *CryptoKey) {
	macro.Rewrite("$_.publicKey = $1", publicKey)
}

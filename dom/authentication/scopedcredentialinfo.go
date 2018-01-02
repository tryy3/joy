package authentication

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/crypto"
)

type ScopedCredentialInfo struct {
}

func (*ScopedCredentialInfo) Credential() (credential *ScopedCredential) {
	macro.Rewrite("$_.credential")
	return credential
}

func (*ScopedCredentialInfo) PublicKey() (publicKey *crypto.CryptoKey) {
	macro.Rewrite("$_.publicKey")
	return publicKey
}

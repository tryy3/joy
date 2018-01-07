package scopedcredentialinfo

import (
	"github.com/matthewmueller/joy/dom/authentication"
	"github.com/matthewmueller/joy/dom/crypto"
	"github.com/matthewmueller/joy/macro"
)

// ScopedCredentialInfo struct
// js:"ScopedCredentialInfo,omit"
type ScopedCredentialInfo struct {
}

// Credential prop
// js:"credential"
func (*ScopedCredentialInfo) Credential() (credential *authentication.ScopedCredential) {
	macro.Rewrite("$_.credential")
	return credential
}

// PublicKey prop
// js:"publicKey"
func (*ScopedCredentialInfo) PublicKey() (publicKey *crypto.CryptoKey) {
	macro.Rewrite("$_.publicKey")
	return publicKey
}

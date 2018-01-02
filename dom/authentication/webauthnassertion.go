package authentication

import (
	"github.com/matthewmueller/joy/macro"
)

type WebAuthnAssertion struct {
}

func (*WebAuthnAssertion) AuthenticatorData() (authenticatorData []byte) {
	macro.Rewrite("$_.authenticatorData")
	return authenticatorData
}

func (*WebAuthnAssertion) ClientData() (clientData []byte) {
	macro.Rewrite("$_.clientData")
	return clientData
}

func (*WebAuthnAssertion) Credential() (credential *ScopedCredential) {
	macro.Rewrite("$_.credential")
	return credential
}

func (*WebAuthnAssertion) Signature() (signature []byte) {
	macro.Rewrite("$_.signature")
	return signature
}

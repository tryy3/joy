package authentication

import (
	"github.com/matthewmueller/joy/macro"
)

type WebAuthentication struct {
}

func (*WebAuthentication) GetAssertion(assertionChallenge []byte, options *AssertionOptions) (w *WebAuthnAssertion) {
	macro.Rewrite("await $_.getAssertion($1, $2)", assertionChallenge, options)
	return w
}

func (*WebAuthentication) MakeCredential(accountInformation *Account, cryptoParameters []*ScopedCredentialParameters, attestationChallenge []byte, options *ScopedCredentialOptions) (s *ScopedCredentialInfo) {
	macro.Rewrite("await $_.makeCredential($1, $2, $3, $4)", accountInformation, cryptoParameters, attestationChallenge, options)
	return s
}
